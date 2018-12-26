---
layout: post
title: "Segment Tree / Index Tree"
slug: "segment-index-tree"
date: 2018-10-12
categories: data-structure
tags: [data structure]
---

# Segment and Index Tree

**세그먼트 트리(Segment Tree)** 또는 **인덱스 트리(Index Tree)**는 보통 다음 연산을 효율적으로 처리하기 위한 자료구조이다.

수열(시퀀스) $\{a_n\}$ 이 주어졌을 때, 다음 연산이 연속적으로 일어난다.

1. 구간 $l$ 부터 $r$ 까지의 합 $a_l + a_{l+1} +  \cdots + a_{r-1} + a_r$  을 구한다. 
2. $i$ 번째 수 $a_i$ 의 값을 $v$ 로 바꾼다. $(a_i = v)$ 

일반적으로, 1번 문제는 누적합을 통해서 쉽게 구할 수 있다.  예를 들어, $S_k$ 를 $k$ 까지의 합이라고 한다면, 구간 $l$ 부터 $r$ 까지의 합은 $S_r -  S_{l-1}$ 이 된다.

하지만, 2번 연산으로 인해 1에서 구한 누적합이 계속해서 변하게 되고 이를 갱신하기 위해서 $O(n)$ 의 시간이 든다. **세그먼트 트리는 시퀀스 자료를 구간별로 나누어 트리에 저장한 것으로, 위 두 연산을 $O(\log n)$ 에 수행할 수 있다.**



## 초기 트리 구성

$N$ 개의 시퀀스형 자료가 있다고 하자. $2^M \geq N$ 을 만족하는 가장 작은 $M$ 에 대하여, $2^M$ 개의 리프 노드로 구성된 **포화 이진 트리(Full Binary Tree)** 를 만든다. 리프 노드는 시퀀스 자료가 순서대로 들어가 있으며, 비어있는 노드는 0으로 채운다. 예를 들어, 배열 [3,5,2,9,1,7] 로 구성한 포화 이진 트리는 아래 그림과 같다. 각 노드에서 노란색 범위는 자료의 구간을 나타내며, 흰 숫자는 구간의 합을 나타낸다.

![초기 트리 구성](https://github.com/sjnov11/sjnov11.github.com/blob/master/_img/2018/10/12/segment_tree.png?raw=true)

### 



## 구간합 구하기 (1번 연산)

위의 예제에서 구간 [2:4] 의 합을 구해보자. 아래 그림에서 빨간색으로 색칠된 노드들의 탐색을 마치면 답을 구할 수 있다.

![구간 합 구하기](https://github.com/sjnov11/sjnov11.github.com/blob/master/_img/2018/10/12/segment_tree_find_sum.png?raw=true)



합을 구할 때, 루트에서부터 탐색을 시작한다. 노드가 갖고있는 범위를 [begin, end] 로, 구하고자 하는 구간을 [left, right] 라고 하자. 이 때, 아래의 4가지 경우가 존재한다.

1. [begin, end] 가 [left, right] 에 포함되는 경우
2. [begin, end] 가 [left, right] 에 포함되지 않는 경우
3. [begin, end] 가 [left, right] 를 완전히 포함하는 경우
4. [begin, end] 가 [left, right] 와 일부분이 겹치는 경우



1번의 경우, 현재 노드가 구하고자 하는 범위내의 값을 모두 가지고 있는 것이므로 노드의 값을 리턴하고 탐색을 종료한다. 반대로 2 번의 경우는 노드가 구하고자 하는 범위 밖이므로 0을 리턴하고 탐색을 종료한다.

3번,4번의 경우는 자식을 기준으로 탐색을 다시 진행하여 1번,2번 경우가 나올 때 까지 재귀 호출 한다.



## 자료값 변경 (2번 연산)

값을 바꾸면 해당 자료를 구간으로 담고있는 모든 노드들 - 바꾸는 자료값을 담는 리프 노드에서 루트 노드까지의 노드들- 을 업데이트 해주어야 한다. 예제에서 3번 노드의 값을 8로 바꾸었을 때, 변경해야하는 노드들을 빨간색 노드로 나타내면 다음과 같다.



![자료값 변경](https://github.com/sjnov11/sjnov11.github.com/blob/master/_img/2018/10/12/segment_tree_update.png?raw=true)


