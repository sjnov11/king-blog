---
layout: post
title: "Union-Find 와 Graph의 Cycle Detection"
slug: "union-find-cycle-detection"
date: 2018-10-02
categories: algorithm
tags: [algorithm]
---
## Union-Find

Union-Find는 서로소 집합(disjoint set)을 관리하는 자료구조이다. Union-Find는 이름처럼 다음 두 가지 연산을 가지고 있다.

1) **Find(A)**: 요소(element) A가 속한 집합을 반환한다.
2) **Union(A, B)**: 요소 A가 속한 집합과 요소 B가 속한 집합을 합친다.

다음의 예제에서 Find(1) == Find(2) 이고, Union(1, 3) 또는 Union(2, 3) 의 결과는 아래와 같다.

![set1과 set2가 하나로 되었다](https://github.com/sjnov11/sjnov11.github.com/blob/master/_img/2018/08/20/union-find_1.jpg?raw=true)



## Union-Find 구현

배열을 이용한 구현 방법과 트리를 이용한 구현 방법이 있다. 배열로 구현한 경우 Find를 $O(1)$ 시간으로 처리할 수 있지만, Union의 경우는 $O(N)$ 의 시간이 걸린다. 반면, 트리로 구현한 경우 Find와 Union 모두 $O(logN)$ 시간이 걸린다.



### 1) 배열을 이용한 구현 방법

그래프의 각 정점이 어떤 set 에 속하는 지 배열에 담아두는 방법이다. 위 예제 초기 그림에서 set_num[1] = 1, set_num[2] = 1, set_num[3] = 2, set_num[4] = 3 이다.

#### (1) 초기화

Union-Find의 초기 상태는 각 요소들이 모두 서로소 집합인 상태이다. 아래와 같이 요소들이 속하는 집합을 자기 자신으로 두면 된다.

```c++
//code by RiKang, weeklyps.com
struct union_find{
    vector<int> set_num;  
    void init(int n){ 
        set_num.resize(n + 1);
        for(int i = 1; i <= n; i++)
            set_num[i] = i;
    }
}uf;
```

#### (2) Find

Find는 set_num 배열의 index를 참조하여 구할 수 있다. ($O(1)$)

```c++
int find(int elem1) {
    return set_num[elem1];
}
```

#### (3) Union

Union은(elem1을 기준) elem2가 속한 집합의 모든 원소를 elem1의 집합으로 옮기면 된다. ($O(N)$)

```c++
void union (int elem1, int elem2) {
    int remove_set = find(elem2);
    for (int i = 1; i < set_num.size(); i++) {
        if (set_num[i] == remove_set)
        	set_num[i] = find(elem1);
    }
}
```



### 2) 트리를 이용한 구현

배열과는 달리 이 방법에서는 집합의 번호를 루트의 번호로 대신한다. 각 요소들의 부모를 저장해두면 루트까지 타고 올라갈 수 있으므로, 배열에서 요소들이 집합 번호를 담고 있던 것(set_num) 과 달리 요소들이 부모 요소를 찾아갈 수 있도록 parent 배열을 만든다. 이 parent 가 결국 트리의 adj_list를 만드는 것과 같다.

![](https://github.com/sjnov11/sjnov11.github.com/blob/master/_img/2018/08/20/union-find_2.jpg?raw=true)

 

#### (1) 초기화

배열을 이용한 구현과 같다. 초기 상태에서 부모가 없으므로 자기 자신을 부모로 둔다. 혹은 -1을 두어도 좋다. 필자는 루트 노드의 부모를 랭크* (-1)로 표현하였다.

```c++
vector<int> parent;
void init(int V) {
    parent.resize(V+1);
    for (int i = 1; i <= V; i++) {
        parent[i] = -1;
    }
}
```

#### (2) Find

parent가 자기자신(혹은 음수)이 나올 때 까지 부모를 계속 탐색한다. ($O(\log{N})$)

```c++
int find(int e) {
    if (parent[e] < 0) return e;
    return find(parent[e]);
}
```

#### (3) Union

Union은(elem1을 기준) elem2의 루트를 elem1의 루트의 자식으로 달아주면 된다. ($O(\log{N})$) 

```c++
void uni (int e1, int e2) {
    int r1 = find(e1);
    int r2 = find(e2);
    parent[r2] = r1;
}
```

Union시 최악의 경우 LinkedList 형태처럼 줄줄이 달릴 수가 있다. 두 집합 중 높이가 낮은 트리를 높은 트리 아래에 넣어두면 최악의 경우를 피할 수 있다. 

 ```c++
void uni (int e1, int e2) {
    int r1 = find(e1);
    int r2 = find(e2);
    int rank_r1 = parent[r1] * (-1);
    int rank_r2 = parent[r2] * (-1);
    
    if (rank_r1 > rank_r2) 
    	parent[r2] = r1;
    else {
        parent[r1] = r2;
        if (rank_r1 == rank_r2) 
        	parent[r2]--;
    }
}
 ```



## Python 코드

```python
class DisjointSet_list:
    def __init__(self, V):
        self.V = V
        self.set_list = [v for v in range(self.V + 1)] 
        
    def find(self, e):
        return self.set_list[e]

    def union(self, e1, e2):
        remove_set_num = self.find(e2)
        for v, set_num in enumerate(self.set_list):
            if set_num == remove_set_num:
                self.set_list[v] = e1


class DisjointSet_tree:
    def __init__(self, V):
        self.V = V
        self.parent_list = [-1 for v in range(V+1)]

    def find(self, e):
        parent = self.parent_list[e]
        if parent < 0:
            return e
        return self.find(parent)

    def union(self, e1, e2):
        r1 = self.find(e1)
        r2 = self.find(e2)

        if r1 == r2:
            return

        rank_r1 = self.parent_list[r1] * (-1)
        rank_r2 = self.parent_list[r2] * (-1)

        if rank_r1 > rank_r2:
            self.parent_list[r2] = r1
        else:
            self.parent_list[r1] = r2
            if rank_r1 == rank_r2:
                self.parent_list[r2]-=1
```





## Undirected Graph의 Cycle Detection

아무 node도 연결되지 않은 상태를 초기 상태로 두자. 각 node들은 자기 자신으로만 이루어져 있는 그래프가 될 것이다. 각각의 독립된 그래프를 disjoint set으로 보자. edge로 연결하게 되면 두 독립된 그래프가 하나의 그래프가 된다. 즉, 두 집합이 하나의 집합으로 Union 된다. edge로 연결할 때 두 node가 같은 그래프 집합에 있다면, Cycle이 발생하는 것이다. 

```c++
bool isCycle (Graph graph) {
	for (int edge : edge_list) {
		// edge : (u, v)
		int n1 = edge.first;
		int n2 = edge.second;
    	if (find(n1) == find(n2)) return true;
    	union(n1, n2);
	}
	return false;
}
```

