func findOrder(numCourses int, prerequisites [][]int) []int {
    g := NewGraph()
    for i := 0; i < numCourses; i++ {
        g.AddVertex(i)
    }
    for _, prerequisite := range prerequisites {
        to, from := prerequisite[0], prerequisite[1]
        g.AddEdge(from, to)
    }

    var result []int
    for _, data := range g.TopoSort() {
        result = append(result, data.(int))
    }
    if len(result) < numCourses {
        return nil
    }
    return result
}
