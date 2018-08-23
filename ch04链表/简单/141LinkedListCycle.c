#include <stdio.h>
#include <stdbool.h>

struct ListNode
{
    int val;
    struct ListNode *next;
};

bool hasCycle(struct ListNode *head)
{
    if (head == NULL)
    {
        return false;
    }
    struct ListNode *fast, *slow;
    slow = fast = head;
    
    while (slow != NULL && fast->next != NULL)
    {
        slow = slow->next;
        fast = fast->next->next;
        if (slow == fast)
        {
            return true;
        }
    }
    return false;
}

int main(int argc, char const *argv[])
{

    return 0;
}
