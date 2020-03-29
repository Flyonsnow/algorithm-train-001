# 学习笔记
## Deque的push与addFirst与addLast
```java
    /**
             * openjdk11 中 push方法调用的就是addFirst
             * 
             */
              public void push(E e) {
                     addFirst(e);
                }
```

##  PriorityQueue 优先队列
优先队列是可以对插入元素进行排序的有序队列，可以根据自然顺序或者自定义比较器对插入元素进行排序，内部也是由数组实现的。

### 1.add offer 插入元素
add方法添加元素的核心排序是调用下面两个方法进行

```java
if (comparator != null)
            siftUpUsingComparator(k, x, queue, comparator);
        else
            siftUpComparable(k, x, queue);
```

其中siftUpUsingComparator顾名思义是调用自定义的比较器进行的。

```java
private static <T> void siftUpComparable(int k, T x, Object[] es) {
        Comparable<? super T> key = (Comparable<? super T>) x;
        while (k > 0) {
            int parent = (k - 1) >>> 1;
            Object e = es[parent];
            if (key.compareTo((T) e) >= 0)
                break;
            es[k] = e;
            k = parent;
        }
        es[k] = key;
    }

    private static <T> void siftUpUsingComparator(
            int k, T x, Object[] es, Comparator<? super T> cmp) {
        while (k > 0) {
            int parent = (k - 1) >>> 1;
            Object e = es[parent];
            if (cmp.compare(x, (T) e) >= 0)
                break;
            es[k] = e;
            k = parent;
        }
        es[k] = x;
    }
```
大概意思就是，当插入一个新元素时，跟前面元素进行比较，如果比之前的元素大直接break插入，如果小则互相交换值并用交换后的值继续与之前的元素比较，直到最小不用交换。

### 2.remove poll 删除元素

删除，先遍历所有元素，发现元素所在的下标，删除后重新排序。

```java
private void siftDown(int k, E x) {
        if (comparator != null)
            siftDownUsingComparator(k, x, queue, size, comparator);
        else
            siftDownComparable(k, x, queue, size);
    }

    private static <T> void siftDownComparable(int k, T x, Object[] es, int n) {
        // assert n > 0;
        Comparable<? super T> key = (Comparable<? super T>)x;
        int half = n >>> 1;           // loop while a non-leaf
        while (k < half) {
            int child = (k << 1) + 1; // assume left child is least
            Object c = es[child];
            int right = child + 1;
            if (right < n &&
                ((Comparable<? super T>) c).compareTo((T) es[right]) > 0)
                c = es[child = right];
            if (key.compareTo((T) c) <= 0)
                break;
            es[k] = c;
            k = child;
        }
        es[k] = key;
    }

    private static <T> void siftDownUsingComparator(
        int k, T x, Object[] es, int n, Comparator<? super T> cmp) {
        // assert n > 0;
        int half = n >>> 1;
        while (k < half) {
            int child = (k << 1) + 1;
            Object c = es[child];
            int right = child + 1;
            if (right < n && cmp.compare((T) c, (T) es[right]) > 0)
                c = es[child = right];
            if (cmp.compare(x, (T) c) <= 0)
                break;
            es[k] = c;
            k = child;
        }
        es[k] = x;
    }
```

### 3.peek

比较简单，直接取的

```java
public E peek() {    return (E) queue[0];}
```

## 笔记
基础真的很重要啊，有的写法想不到竟然是因为有些基础的东西忘了。学习难度、时间稳步提升，感觉再加几波班就要完不成了 -_-! 
