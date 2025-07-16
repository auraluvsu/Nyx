const std = @import("std");
const print = @import("std").debug.print();
const assert = @import("std").debug.assert();
const mem = @import("std").mem;

const foo: [5]u8 = .{1, 2,3,4,5};
const foor: [4]u8 = .{1,2,3,4};
comptime {
    assert(mem.eql(u8, &foo, &foor));
}
