// IntersectionOfTwoLinkedList.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>

class ListNode
{
public:
    int data;
    ListNode* next;
    ListNode(int d)
        : data(d)
        , next(nullptr)
    {
    }
};

ListNode* GetIntersectionNode(ListNode* A, ListNode* B)
{
    while (A != nullptr)
    {
        ListNode* pB = B;
        while (pB != nullptr)
        {
            if (pB == A)
                return A;
            pB = pB->next;
        }
        A = A->next;
    }

    return nullptr;
}

ListNode* GetListFromInput(ListNode** ppHead, int count)
{
    if (count == 0)
        return nullptr;

    std::cout << "Enter " << count << " numbers of the list: ";
    ListNode* pTail = *ppHead;
    for (int i = 0; i < count; ++i)
    {
        int data = 0;
        std::cin >> data;
        if (*ppHead == nullptr)
        {
            *ppHead = new ListNode(data);
            pTail = *ppHead;
        }
        else
        {
            pTail->next = new ListNode(data);
            pTail = pTail->next;
        }
    }

    return pTail;
}

void FreeList(ListNode* pNode)
{
    if (pNode == nullptr)
        return;

    FreeList(pNode->next);
    delete pNode;
}

int main()
{
    while (true)
    {
        int countA = 0, countB = 0;
        std::cout << "Number of nodes in list A and B: ";
        std::cin >> countA >> countB;
        if (countA == 0 && countB == 0)
            break;

        ListNode* aHead = nullptr, * bHead = nullptr;
        ListNode* aTail = GetListFromInput(&aHead, countA);
        ListNode* bTail = GetListFromInput(&bHead, countB);

        std::cout << "Number of common nodes: ";
        int countC = 0;
        std::cin >> countC;
        ListNode* cHead = nullptr;
        GetListFromInput(&cHead, countC);

        if (aHead == nullptr)
            aHead = cHead;
        else
            aTail->next = cHead;

        if (bHead == nullptr)
            bHead = cHead;
        else
            bTail->next = cHead;

        ListNode* intersectionNode = GetIntersectionNode(aHead, bHead);
        std::cout << "The intersection node is: " << (intersectionNode != nullptr ? intersectionNode->data : -1) << std::endl << std::endl;
    
        if (cHead != nullptr)
        {
            aTail->next = nullptr;
            bTail->next = nullptr;
        }
        FreeList(aHead);
        FreeList(bHead);
        FreeList(cHead);

    }
}
