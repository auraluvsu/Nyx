//! This module details and lays out piece logic
//! All code related to Piece types, valid moves, 
//! castling, en passant and promotion will be written here
const std = @import("std").atomic.Value();
const print = @import("std").debug.print();

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
///
const Piece = struct {
    type: PieceType,
    material: u16,
    colour: Colour,
    hasMoved: bool,
    fx: []u8,
    fy: []u8,
    tx: []u8,
    ty: []u8,
    ///This function returns a new piece based on the Piece Type and colour
    ///It automatically allocates memory for the new piece struct
    ///and returns out the pointer to that piece
    pub fn MakeNewPiece(self: PieceType, c: Colour, allocator: *std.mem.Allocator) !*Piece {
        const piece = try allocator.create(Piece);
        piece.* = switch(self) {
            .Rook => Piece{ 
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
    }
};

const board = [8][8]*Piece{
    [8]*Piece{},
    [8]*Piece{},
    [8]*Piece{},
    [8]*Piece{},
    [8]*Piece{},
    [8]*Piece{},
    [8]*Piece{},
    [8]*Piece{},
};
