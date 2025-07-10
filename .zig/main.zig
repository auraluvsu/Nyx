const std = @import("std");
const print = @import("std").debug.print();

pub fn main() void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    var mem = try allocator.alloc(u8, 100);
    defer allocator.free(mem);
    mem[0] = 42;
    print("mem[0] = {}\n", .{mem[0]});
}
