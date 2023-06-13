import math
import random


class Sample(object):
    def __init__(self, values: list[int], weights: list[float]):
        self.values = values
        self.weights = weights

        s = sum(self.weights)
        self.weights = [x / float(s) if s != 0 else 0.0 for x in self.weights]

        self.cum_sum = [0] * len(self.values)
        self.compute_cum_sum()

    def compute_cum_sum(self):
        cum_sum = [0] * len(self.values)
        s = 0
        for i in range(len(self.values)):
            s += self.weights[i]
            cum_sum[i] = s

        self.cum_sum = cum_sum

    def get_one(self):
        u = random.uniform(0, 1)

        left, right = 0, len(self.values) - 1
        last_true = 0
        while left <= right:
            mid = int((left + right) / 2)
            if self.cum_sum[mid] >= u:
                last_true = mid
                right = mid - 1
            else:
                left = mid + 1

        return self.values[last_true]


class Node(object):
    def __init__(self, key=None, val=None, level=None):
        self.key = key
        self.val = val
        self.next_pointers = {}
        self.prev_pointers = {}
        self.level = level


class SkipList(object):
    def __init__(self, max_size=10 ** 6):
        self.levels = range(1, int(math.log(max_size, 2)) + 2)

        self.probabilities = [1.0] * len(self.levels)
        for i in range(len(self.levels)):
            self.probabilities[i] = 0.5 * self.probabilities[i - 1] if i > 0 else 1.0

        self.sample = Sample(self.levels, self.probabilities)
        node = Node('__head__', -float("Inf"), len(self.levels))

        for i in self.levels:
            node.next_pointers[i] = None
            node.prev_pointers[i] = None

        self.head = node
        self.tail = node

        self.node_map = {}

    def search_key(self, key):
        return self.node_map[key].val if key in self.node_map else False

    def search_val(self, val):
        curr_level = self.levels[-1]
        node = self.head

        for i in range(curr_level, 0, -1):
            while node.next_pointers[i] is not None and node.next_pointers[i].val < val:
                node = node.next_pointers[i]

        return node.next_pointers[1] is not None and node.next_pointers[1].val == val

    def insert(self, key, val):
        if key not in self.node_map:
            curr_level = self.levels[-1]
            node = self.head

            random_level = self.sample.get_one()
            updates = [None] * random_level

            for i in range(curr_level, 0, -1):
                while node.next_pointers[i] is not None and node.next_pointers[i].val < val:
                    node = node.next_pointers[i]

                if i <= random_level:
                    updates[i - 1] = node

            new_node = Node(key, val, random_level)

            for i in range(1, random_level + 1):
                node = updates[i - 1]

                new_node.next_pointers[i] = node.next_pointers[i]

                if node.next_pointers[i] is not None:
                    node.next_pointers[i].prev_pointers[i] = new_node
                else:
                    if i == 1:
                        self.tail = new_node

                node.next_pointers[i] = new_node
                new_node.prev_pointers[i] = node

            self.node_map[key] = new_node

        else:
            self.delete(key)
            self.insert(key, val)

    def delete(self, key):
        if key in self.node_map:
            node = self.node_map[key]

            for i in range(1, node.level + 1):
                if node.prev_pointers[i] is not None:
                    node.prev_pointers[i].next_pointers[i] = node.next_pointers[i]

                if node.next_pointers[i] is not None:
                    node.next_pointers[i].prev_pointers[i] = node.prev_pointers[i]
                else:
                    if i == 1:
                        self.tail = node.prev_pointers[i]

            self.node_map.pop(key)

    def get_min(self):
        if self.head is not None and self.head.next_pointers[1] is not None:
            return self.head.next_pointers[1].key, self.head.next_pointers[1].val
        return None

    def get_max(self):
        if self.tail is not None:
            return self.tail.key, self.tail.val
        return None

    def pop_min(self):
        out = self.get_min()
        if out is not None:
            self.delete(out[0])
        return out

    def pop_max(self):
        out = self.get_max()
        if out is not None:
            self.delete(out[0])
        return out

    def get_predecessor(self, val):
        curr_level = self.levels[-1]
        node = self.head

        for i in range(curr_level, 0, -1):
            while node.next_pointers[i] is not None and node.next_pointers[i].val < val:
                node = node.next_pointers[i]

        if node is not None:
            return node.key, node.val

        return None

    def get_successor(self, val):
        curr_level = self.levels[-1]
        node = self.head

        for i in range(curr_level, 0, -1):
            while node.next_pointers[i] is not None and node.next_pointers[i].val < val:
                node = node.next_pointers[i]

        if node.next_pointers[1] is not None:
            if node.next_pointers[1].val > val:
                return node.next_pointers[1].key, node.next_pointers[1].val
            else:
                if node.next_pointers[1].next_pointers[1] is not None:
                    return node.next_pointers[1].next_pointers[1].key, node.next_pointers[1].next_pointers[1].val
        return None

    def print_list(self):
        for i in range(len(self.levels) - 1, -1, -1):
            p = self.head
            out = []

            while p is not None:
                out.append((p.key, p.val))
                p = p.next_pointers[self.levels[i]]

            print(out)


sl = SkipList()

sl.insert(3, 'hello')
sl.insert(11, 'world')
sl.insert(18, 'this')
sl.insert(9, 'is')
sl.insert(14, 'me')
sl.insert(19, 'reaching')
sl.insert(17, 'out')
sl.insert(6, 'to')
sl.insert(24, 'you')
sl.insert(27, 'to')
sl.insert(10, 'talk')
sl.insert(2, 'about')
sl.insert(4, 'the')
sl.insert(26, 'extended')
sl.insert(29, 'warranty')
sl.insert(25, 'of')
sl.insert(23, 'your')
sl.insert(21, 'car')
sl.insert(22, 'insaurance')
sl.insert(20, 'please')
sl.insert(5, 'response')
sl.insert(12, 'ASAP')
sl.insert(13, 'because')
sl.insert(1, 'this')
sl.insert(28, 'is')
sl.insert(16, 'a')
sl.insert(8, 'limited')
sl.insert(15, 'time')
sl.insert(7, 'offer')

print(sl.search_val('hello'))
# sl.delete
print(sl.search_val('world'))
print(sl.search_val('this'))
print(sl.search_val('is'))
print(sl.search_val('me'))
print(sl.search_val('reaching'))
print(sl.search_val('out'))
print(sl.search_val('to'))
print(sl.search_val('you'))
print(sl.search_val('to'))
print(sl.search_val('talk'))
print(sl.search_val('about'))
print(sl.search_val('the'))
print(sl.search_val('extended'))
print(sl.search_val('warranty'))
print(sl.search_val('of'))
print(sl.search_val('your'))
print(sl.search_val('car'))
print(sl.search_val('insaurance'))
print(sl.search_val('please'))
print(sl.search_val('response'))
print(sl.search_val('ASAP'))
print(sl.search_val('because'))
print(sl.search_val('this'))
print(sl.search_val('is'))
print(sl.search_val('a'))
print(sl.search_val('limited'))
print(sl.search_val('time'))
print(sl.search_val('offer'))

sl.print_list()
