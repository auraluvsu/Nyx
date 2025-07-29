const piece = @import("piece.zig");
const std = @import("std");

pub fn foo() void {
    const allocator = std.mem.Allocator;
    const newPiece = piece.PieceType.Rook;
    piece.printBoard();
    const rook = piece.MakeNewPiece(newPiece, piece.Colour.White, allocator);
}
