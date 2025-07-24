//! This module details and lays out piece logic
//! All code related to Piece types, valid moves, 
//! castling, en passant and promotion will be written here
const std = @import("std");
const print = @import("std").debug.print();
const thread = @import("std").Thread;
const io = @import("io");
const os = @import("os");

const PieceType = enum {
    Pawn,
    Rook,
    Bishop,
    Knight,
    Queen,
    King,
};

const Colour = enum {
    White,
    Black,
};

/// Each piece has coordinates: from X, from Y, to X and to Y
/// Each piece also has a Piece type
/// Pieces also have a meterial count according to their piece type
/// Pieces have a hasMoved section specifically for castling logic
const Piece = struct {
    type: PieceType,
    material: u16,
    colour: Colour,
    hasMoved: bool,
    fx: []u8,
    fy: []u8,
    tx: []u8,
    ty: []u8,
};
///This function returns a new piece based on the Piece Type and colour
///It automatically allocates memory for the new piece struct
///and returns out the pointer to that piece
pub fn MakeNewPiece(self: PieceType, c: Colour, allocator: *std.mem.Allocator) !*Piece {
    const piece = try allocator.create(Piece);
    piece.* = switch(self) {
        .Rook => Piece {
            .type = PieceType.Rook,
            .material = 500,
            .colour = c,
            .hasMoved = false,
        },
        .Knight => Piece{
            .type = PieceType.Knight,
            .material = 300,
            .colour = c,
            .hasMoved = false,
        },
        .Bishop => Piece{
            .type = PieceType.Bishop,
            .material = 300,
            .colour = c,
            .hasMoved = false,
        },
        .Queen => Piece{
            .type = PieceType.Queen,
            .material = 900,
            .colour = c,
            .hasMoved = false,
        },
        .King => Piece{
            .type = PieceType.King,
            .material = 0,
            .colour = c,
            .hasMoved = false,
        },
        .Pawn => Piece{
            .type = PieceType.Pawn,
            .material = 100,
            .colour = c,
            .hasMoved = false,
        },
    };
    return piece;
}



///This is the board struct that will be used for calculating moves 
///and printing out the board to the terminal
pub fn printBoard() [8][8]*Piece {
    const board = [8][8]*Piece{};
    const order = []PieceType{
        PieceType.Rook,
        PieceType.Knight,
        PieceType.Bishop,
        PieceType.Queen,
        PieceType.King,
        PieceType.Bishop,
        PieceType.Knight,
        PieceType.Rook,
    };
    for (0..8) |i| {
        board[i][0] = Piece.MakeNewPiece(order, Colour.Black, *std.mem.Allocator);
        board[i][1] = Piece.MakeNewPiece(PieceType.Pawn, Colour.Black, *std.mem.Allocator);
        board[i][6] = Piece.MakeNewPiece(PieceType.Pawn, Colour.White, *std.mem.Allocator);
        board[i][7] = Piece.MakeNewPiece(order, Colour.White, *std.mem.Allocator);
    }
    return board;
}
