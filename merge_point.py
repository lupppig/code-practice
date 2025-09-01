#!/bin/python3

import math
import os
import random
import re
import sys


class SinglyLinkedListNode:
    def __init__(self, node_data):
        self.data = node_data
        self.next = None


class SinglyLinkedList:
    def __init__(self):
        self.head = None
        self.tail = None

    def insert_node(self, node_data):
        node = SinglyLinkedListNode(node_data)

        if not self.head:
            self.head = node
        else:
            self.tail.next = node

        self.tail = node


def print_singly_linked_list(node, sep, fptr):
    while node:
        fptr.write(str(node.data))

        node = node.next

        if node:
            fptr.write(sep)


# Complete the findMergeNode function below.


#
# For your reference:
#
# SinglyLinkedListNode:
#     int data
#     SinglyLinkedListNode next
#
#
def findMergeNode(head1, head2):
    h1 = head1
    h2 = head2

    ll_of_h1, ll_of_h2 = 0, 0

    while h1 is not None or h2 is not None:
        if h1 is not None:
            h1 = h1.next
            ll_of_h1 += 1

        if h2 is not None:
            h2 = h2.next
            ll_of_h2 += 1

    diff_len = abs(ll_of_h1 - ll_of_h2)

    t_node = head1 if ll_of_h1 >= ll_of_h2 else head2
    t_node1 = head2 if ll_of_h2 <= ll_of_h1 else head1

    while t_node.next is not None and diff_len != 0:
        diff_len -= 1
        t_node = t_node.next

    while t_node1 is not None and t_node is not None:
        if t_node1 is t_node:
            return t_node.data
        t_node = t_node.next
        t_node1 = t_node1.next


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")

    tests = int(input())

    for tests_itr in range(tests):
        index = int(input())

        llist1_count = int(input())

        llist1 = SinglyLinkedList()

        for _ in range(llist1_count):
            llist1_item = int(input())
            llist1.insert_node(llist1_item)

        llist2_count = int(input())

        llist2 = SinglyLinkedList()

        for _ in range(llist2_count):
            llist2_item = int(input())
            llist2.insert_node(llist2_item)

        ptr1 = llist1.head
        ptr2 = llist2.head
        for i in range(llist1_count):
            if i < index:
                ptr1 = ptr1.next

        for i in range(llist2_count):
            if i != llist2_count - 1:
                ptr2 = ptr2.next

        ptr2.next = ptr1

        result = findMergeNode(llist1.head, llist2.head)

        fptr.write(str(result) + "\n")

    fptr.close()
