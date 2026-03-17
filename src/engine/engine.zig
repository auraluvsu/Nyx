const std  = @import("std");

const value = union(enum) {
    str: []const u8,
    num: i32,
};
