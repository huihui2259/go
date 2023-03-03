package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head.Next == nil || k == 1 {
		return head
	}
	dummy := &ListNode{}
	preHead, Head, Tail, proTail := dummy, head, head, head.Next
	for {
		i := 1
		for i < k && Tail != nil {
			Tail = Tail.Next
			i++
		}
		if Tail == nil {
			break
		}
		proTail = Tail.Next
		Tail.Next = nil
		Tail = reverse(Head)
		preHead.Next = Tail
		Head.Next = proTail

		preHead = Head
		Head, Tail = proTail, proTail
	}
	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	cur := pre.Next
	for cur != nil {
		nextCur := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextCur
	}
	head.Next = nil
	return pre
}
