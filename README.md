# protoc-gen-sql
This is a `protoc` plugin that generates `.sql` files from `.proto` for database migration.
It has to be used along with `protoc`.

## Install
In the `protoc-gen-sql`
```shell
go install .
```
It will install the plugin in your `$GOBIN`

## Run the plugin

```
protoc --sql_out=$DST_DIR --sql_opt=file=eva $SRC_DIR/xxxxx.proto
```

It will generate a file named `eva.sql` in your destination directory.

You can pass the `file` to `--sql_opt` to specify the name of generated file, otherwise it will use the default name
same as `.proto` file.

For example, if you have a `.proto`:

```protobuf
syntax = "proto3";
package compound;

import "google/protobuf/timestamp.proto";

option go_package = "rukolahasser/test";

message MyMessage {
  google.protobuf.Timestamp ts = 1;
  // index
  uint64 block = 2;
  bytes idx = 3; // uint256
  repeated int32 arr = 6;
}
```

It will generate a table:
```sql
CREATE TABLE IF NOT EXISTS "rukolahasser.test.mymessage"
(
    "ts"    timestamptz NOT NULL,
    "block" int8,
    "idx"   NUMERIC(78, 0),
    "arr"   int8[]
);
CREATE INDEX ON "rukolahasser.test.mymessage" (block);
```

For more usage information, please refer to `protoc` docs.

# Leading Comments

You may want to create an index column, you can use leading comments right above the field:
```protobuf
syntax = "proto3";

message MyMessage {
  // index
  int32 indexed_block = 1;
}
```

## Trailing Comments

Sometimes you want to generate different types in postgres, you can use trailing comments:

```protobuf
syntax = "proto3";

message MyMessage {
  int32 block = 1;
  bytes idx = 2; // uint256
}
```

So it will generate a table which has a column `"idx" NUMERIC(78,0)` instead of `"idx" bytea`.

Trailing comment and postgres type mapping:

| comment type   | postgres type |
|----------------|---------------|
| int256/uint256 | NUMERIC(78,0) |
| int248/uint248 | NUMERIC(75,0) |
| int240/uint240 | NUMERIC(73,0) |
| int232/uint232 | NUMERIC(70,0) |
| int224/uint224 | NUMERIC(68,0) |
| int216/uint216 | NUMERIC(66,0) |
| int208/uint208 | NUMERIC(63,0) |
| int200/uint200 | NUMERIC(61,0) |
| int192/uint192 | NUMERIC(58,0) |
| int184/uint184 | NUMERIC(56,0) |
| int176/uint176 | NUMERIC(53,0) |
| int168/uint168 | NUMERIC(51,0) |
| int160/uint160 | NUMERIC(49,0) |
| int152/uint152 | NUMERIC(46,0) |
| int144/uint144 | NUMERIC(44,0) |
| int136/uint136 | NUMERIC(41,0) |
| int128/uint128 | NUMERIC(39,0) |
| int120/uint120 | NUMERIC(37,0) |
| int112/uint112 | NUMERIC(34,0) |
| int104/uint104 | NUMERIC(32,0) |
| int96/uint96   | NUMERIC(29,0) |
| int88/uint88   | NUMERIC(27,0) |
| int80/uint80   | NUMERIC(25,0) |
| int72/uint72   | NUMERIC(22,0) |
| int64/uint64   | INT8          |
| int56/uint56   | INT8          |
| int48/uint48   | INT8          |
| int40/uint40   | INT8          |
| uint32         | INT8          |
| int32          | INT4          |
| int24/uint24   | INT4          |
| uint16         | INT4          |
| int16          | INT2          |
| int8/uint8     | INT2          |

