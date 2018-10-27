/*
Class Solution takes two numbers represented by linked lists and adds them together, as per LeetCode.com problem #2.

Written by Keith Drew, 2018.
*/
#include <iostream>
#include <string>

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

void print_node(ListNode *ln);
void print_list(ListNode *list);

class Solution {
    public:
        ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
            if (l1 == NULL || l2 == NULL) {
                return NULL;
            }
            
            int carry = 0;
            if (l1->val + l2->val >= 10) {
                carry = 1;
            }

            ListNode *tmpNode = new ListNode((l1->val + l2->val) % 10);
            l1 = l1->next;
            l2 = l2->next;
            ListNode *newList = tmpNode;
            
            while (l1 != NULL && l2 != NULL) {
                tmpNode->next = new ListNode((l1->val + l2->val + carry) % 10);
                tmpNode = tmpNode->next;

                if (l1->val + l2->val + carry >= 10) {
                    carry = 1;
                } else {
                    carry = 0;
                }

                l1 = l1->next;
                l2 = l2->next;
            }

            while (l1 != NULL) {
                tmpNode->next = new ListNode((l1->val + carry) % 10);
                tmpNode = tmpNode->next;

                if (l1->val + carry >= 10) {
                    carry = 1;
                } else {
                    carry = 0;
                }

                l1 = l1->next;
            }

            while (l2 != NULL) {
                tmpNode->next = new ListNode((l2->val + carry) % 10);
                tmpNode = tmpNode->next;

                if (l2->val + carry >= 10) {
                    carry = 1;
                } else {
                    carry = 0;
                }

                l2 = l2->next;
            }

            if (carry == 1) {
                tmpNode->next = new ListNode(1);
            }

            return newList;
        }
};

int main() {
    Solution solution;

    ListNode *l1 = new ListNode(2);
    l1->next = new ListNode(4);
    l1->next->next = new ListNode(3);

    ListNode *l2 = new ListNode(5);
    l2->next = new ListNode(6);
    l2->next->next = new ListNode(4);

    print_list(l1);
    print_list(l2);

    ListNode *list = solution.addTwoNumbers(l1, l2);

    print_list(list);

    ListNode *tmp = list;

    std::cout << std::endl;

    return 0;
}

void print_node(ListNode *ln) {
    if (ln != NULL) {
        std::cout << ln->val << std::endl;
    }
}
void print_list(ListNode *list) {
    ListNode *tmp = list;
    while (tmp != NULL) {
        std::cout << tmp->val << " ";
        tmp = tmp->next;
    }
    std::cout << std::endl;
}