  - genMoves always calls p.IsValidMove(..., nil) so the pawn validator  
    never sees an en-passant target square; IsValidPawnMove therefore rejects
    every en-passant capture at nyx/engine/eval.go:68-108 and nyx/logic/
    movement.go:78-103, so those moves never reach the perft counter.
  - Even if the move slipped through, Perft just swaps newBoard[tx][ty],
    newBoard[fx][fy] = p, nil without clearing the pawn that should be
    captured from the adjacent square, nor does it propagate a new en-passant
    square down the recursion (nyx/engine/perft.go:13-60). That leaves the
    search tree in an illegal state and silently drops any future en-passant
    responses.
  - While you are in there, note that the promotion branch inside genMoves
    yields the move twice, double-counting every promotion and skewing the
    table (nyx/engine/eval.go:88-100).

  1. Thread an *nyx.Position (or your Undo helper from nyx/engine/move.go)
     through Perft/genMoves so each node carries the current en-passant
     target, flag en-passant captures in the callback, and remove the
     captured pawn when you apply the move.
  2. Fix the promotion branch to emit each promotion once, then rerun the
     standard perft suites (e.g. Kiwipete, position 3) to confirm the counts
     match the chess programming wiki baselines.
