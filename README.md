# Binary Tree Right Side View

> NeetCode 150 · Trees — the left view's one-line mirror: DFS right-first, and the first node at each new depth is the visible one. Go, stdlib only, tested.

## Problem

Standing to the tree's right, which node is visible per level? The rightmost of each — which can live in the *left* subtree when the right side is shallow (pinned by a test).

## Approach

Identical machinery to the left view (dsa-109) with the child calls swapped: traverse root → right → left carrying depth, and `depth == len(view)` fires exactly once per level, on the rightmost node — because right-first ordering guarantees the rightmost arrives at each depth first. The deliberate pairing with dsa-109 makes the parameterization visible: the traversal order *is* the view direction, everything else is shared. The BFS take (last node per level) is equally valid; the DFS one is four lines.

**Complexity:** O(n) time, O(h) recursion.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
