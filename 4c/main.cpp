#include <iostream>
#include <string>
#include <vector>
using namespace std;

typedef struct Node {
    int cnt;
    Node* chlds[26];
}* PNode;

PNode NewNode() {
    PNode node = new(Node);
    node->cnt = 0;
    return node;
}

vector<PNode> free_list;

int create_tree(PNode node, string name) {
    for (int i = 0; i < name.length(); i++) {
        int id = name[i] - 'a';
        if (NULL == node->chlds[id]) {
            node->chlds[id] = NewNode();
            free_list.push_back(node->chlds[id]);
        }
        node = node->chlds[id];
    }
    node->cnt += 1;

    return node->cnt;
}

void free_tree() {
    for (int i = 0; i < free_list.size(); i++) {
        for (int j = 0; j < 26; j++) {
            free_list[i]->chlds[j] = NULL;
        }
        delete(free_list[i]);
    }
}


int main() {
    int n, v;
    cin >> n;

    int id = 0;
    PNode root = NewNode();
    free_list.push_back(root);

    string name;
    for (int i= 0; i < n; i++) {
        cin >> name;
        v = create_tree(root, name);
        if (v == 1) {
            cout << "OK" << endl;
        } else {
            cout << name << v - 1 << endl;
        }
    }
    free_tree();

    return 0;
}
