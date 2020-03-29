package liuzhan;

/**
 * message
 * Created by liuzhan
 * on 2020/3/29 8:15 下午
 */
public class MyCircularDeque {

    int[] myQueue;
    //指向队列头部第 1 个有效数据的位置
    int front;
    //指向队列尾部（即最后 1 个有效数据）的下一个位置，即下一个从队尾入队元素的位置。
    int tail;
    //队列容量
    int capacity;
    /**
     *  因为有循环的出现，要特别注意处理数组下标可能越界的情况。
     *
     * （1）指针后移的时候，索引 + 1，要取模；
     *
     * （2）指针前移的时候，为了循环到数组的末尾，需要先加上数组的长度，然后再对数组长度取模。
     */

    /** Initialize your data structure here. Set the size of the deque to be k. */
    public MyCircularDeque(int k) {
        this.myQueue = new int[k];
        this.front = 0;
        this.tail = 0;
        //为了避免“队列为空”和“队列为满”的判别条件冲突，我们有意浪费了一个位置
        //浪费一个位置是指：循环数组中任何时刻一定至少有一个位置不存放有效元素。
        this.capacity = k+1;
    }

    /** Adds an item at the front of Deque. Return true if the operation is successful. */
    public boolean insertFront(int value) {
        if(isFull()){
            return false;
        }
        // 头部指向第 1 个存放元素的位置
        // 插入时，先减，再赋值
        // 删除时，索引 +1（注意取模）
        front = (front - 1 + capacity) % capacity;
        myQueue[front] = value;
        return true;
    }

    /** Adds an item at the rear of Deque. Return true if the operation is successful. */
    public boolean insertLast(int value) {
        if(isFull()){
            return false;
        }
        // 尾部指向下一个插入元素的位置
        // 插入时，先赋值，再加
        // 删除时，索引 -1（注意取模）
        myQueue[tail] = value;
        tail = (tail+1)%capacity;
        return true;
    }

    /** Deletes an item from the front of Deque. Return true if the operation is successful. */
    public boolean deleteFront() {
        if(isEmpty()){
            return false;
        }
        front = (front + 1 ) % capacity;
        return true;
    }

    /** Deletes an item from the rear of Deque. Return true if the operation is successful. */
    public boolean deleteLast() {
        if(isEmpty()){
            return false;
        }
        tail = (tail-1+capacity)%capacity;
        return true;
    }

    /** Get the front item from the deque. */
    public int getFront() {
        if(isEmpty()){
            return -1;
        }
        return myQueue[front];
    }

    /** Get the last item from the deque. */
    public int getRear() {
        if(isEmpty()){
            return -1;
        }
        return myQueue[tail-1 + capacity]%capacity;
    }

    /** Checks whether the circular deque is empty or not. */
    public boolean isEmpty() {
        return front == tail;
    }

    /** Checks whether the circular deque is full or not. */
    public boolean isFull() {
        //可以这样理解，当 tail 循环到数组的前面，要从后面追上 front，还差一格的时候，判定队列为满。
        return (tail + 1)% capacity == front;
    }
}
