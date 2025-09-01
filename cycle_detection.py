def has_cycle(head):
    slow = head
    fast = head

    while fast.next is not None and fast is not None:
        slow = slow.next
        fast = fast.next

        if fast == slow:
            return 1
    return 0
