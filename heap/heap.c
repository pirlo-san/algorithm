#define HEAP_MAX_SIZE 1000

typedef struct {
	int data[HEAP_MAX_SIZE];
	int size;
	int min; // 1: min heap, 0: max heap
} Heap;

void InitHeap(Heap *heap, int min)
{
	if (heap == NULL)
		return;
	memset(heap, 0, sizeof(*heap));
	heap->min = min;
}

int IsEmpty(Heap *heap)
{
	return heap == NULL || heap->size <= 0;
}

int IsFull(Heap *heap)
{
	return heap != NULL && heap->size >= HEAP_MAX_SIZE;
}

int ReversedWithFather(Heap *heap, int idx)
{
	if (heap == NULL || idx >= heap->size || idx <= 0)
		return 0;

	int parent = (idx - 1) / 2;
	if (heap->min)
		return heap->data[idx] < heap->data[parent];
	return heap->data[idx] > heap->data[parent];
}

void Swap(Heap *heap, int i, int j)
{
	if (heap == NULL)
		return;
	if (i < 0 || i >= heap->size)
		return;
	if (j < 0 || j >= heap->size)
		return;

	int tmp = heap->data[i];
	heap->data[i] = heap->data[j];
	heap->data[j] = tmp;
}

int AdjustUp(Heap *heap, int idx)
{
	if (heap == NULL || idx >= heap->size)
		return -1;
	
	while (idx > 0) {
		int father = (idx - 1) / 2;
	
		if (!ReversedWithFather(heap, idx))
			break;
		Swap(heap, idx, father);
		idx = father;
	}
	return 0;
}

int GetCmpChild(Heap *heap, int idx)
{
	if (heap == NULL || idx < 0 || idx >= heap->size)
		return -1;

	int left = 2 * idx + 1;
	int right = 2 * idx + 2;
	if (left >= heap->size)
		return -1;
	if (right >= heap->size)
		return left;
	if (heap->min)
		return heap->data[left] < heap->data[right] ? left : right;
	return heap->data[left] > heap->data[right] ? left : right;
}

int AdjustDown(Heap *heap, int idx)
{
	if (IsEmpty(heap) || idx < 0 || idx >= heap->size) {
		return -1;
	}

	while (idx < heap->size) {
		int left = 2 * idx + 1;
		int right = 2 * idx + 2;
		int cmpChild = GetCmpChild(heap, idx);
		
		if (cmpChild < 0)
			break;
		if (!ReversedWithFather(heap, cmpChild)) 
			break;

		Swap(heap, idx, cmpChild);
		idx = cmpChild;
	}
	return 0;
}

int AddHeap(Heap *heap, int val)
{
	if (IsFull(heap)) {
		return -1;
	}

	heap->data[heap->size] = val;
	return AdjustUp(heap, heap->size++);
}

int DelHeap(Heap *heap, int *val)
{
	if (IsEmpty(heap)) {
		return -1;
	}
	
	*val = heap->data[0];
	heap->data[0] = heap->data[--heap->size];
	return AdjustDown(heap, 0);
}

int ConstructHeap(int data[], int size, int min, Heap *heap)
{
	if (data == NULL || heap == NULL || size <= 0 || size >= HEAP_MAX_SIZE)
		return -1;

	InitHeap(heap, min);
	for (int i = 0; i < size; i++) 
		AddHeap(heap, data[i]);
	return 0;
}

// LeetCode 1046
int lastStoneWeight(int* stones, int stonesSize){
    Heap maxHeap;
    if (ConstructHeap(stones, stonesSize, 0, &maxHeap) < 0)
        return 0;
    
    while (maxHeap.size > 1) {
        int stone1, stone2;
        DelHeap(&maxHeap, &stone1);
        DelHeap(&maxHeap, &stone2);
        if (stone1 == stone2)
            continue;
        int stone = stone1 - stone2;
        if (stone < 0)
            stone = -stone;
        AddHeap(&maxHeap, stone);
    }

    if (maxHeap.size == 0)
        return 0;
    int stone;
    DelHeap(&maxHeap, &stone);
    return stone;
}