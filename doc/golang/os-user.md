## Overview

Package user allows user account lookups by name or id.

For most Unix systems, this package has two internal implementations of resolving user and group ids to names. One is written in pure Go and parses /etc/passwd and /etc/group. The other is cgo-based and relies on the standard C library (libc) routines such as getpwuid_r and getgrnam_r.

When cgo is available, cgo-based (libc-backed) code is used by default. This can be overridden by using osusergo build tag, which enforces the pure Go implementation.

包user允许通过名字或id查询用户账户。

对于大多数Unix系统，这个包有两个内部实现，将用户和组的ID解析为名字。一个是用纯Go编写的，解析/etc/passwd和/etc/group。另一个是基于cgo的，依赖于标准C库（libc）的例程，如getpwuid_r和getgrnam_r。

当cgo可用时，默认使用基于cgo的（libc支持的）代码。这可以通过使用osusergo build标签来重写，它可以强制执行纯Go实现。

## type Group

```
type Group struct {
	Gid  string // group ID
	Name string // group name
}
```
Group represents a grouping of users.

On POSIX systems Gid contains a decimal number representing the group ID.

组代表用户的分组。

在POSIX系统中，Gid包含一个代表组ID的十进制数字。

### func LookupGroup

`func LookupGroup(name string) (*Group, error)`

LookupGroup looks up a group by name. If the group cannot be found, the returned error is of type UnknownGroupError.

LookupGroup按名称查找一个组。如果找不到该组，返回的错误类型为未知组错误。

### func LookupGroupId

`func LookupGroupId(gid string) (*Group, error)`

LookupGroupId looks up a group by groupid. If the group cannot be found, the returned error is of type UnknownGroupIdError.

LookupGroupId通过groupid查找一个组。如果找不到该组，返回的错误类型是UnknownGroupIdError。

## type UnknownGroupError 

```
type UnknownGroupError string
```
UnknownGroupError is returned by LookupGroup when a group cannot be found.

当无法找到一个组时，LookupGroup会返回UnknownGroupError。

### func (UnknownGroupError) Error

`func (e UnknownGroupError) Error() string`

## type UnknownGroupIdError

```
type UnknownGroupIdError string
```
UnknownGroupIdError is returned by LookupGroupId when a group cannot be found.

当无法找到一个组时，LookupGroupId会返回UnknownGroupIdError。

### func (UnknownGroupIdError) Error 

`func (e UnknownGroupIdError) Error() string`

## type UnknownUserError 

```
type UnknownUserError string
```
UnknownUserError is returned by Lookup when a user cannot be found.

当无法找到一个用户时，Lookup会返回UnknownUserError。

### func (UnknownUserError) Error

`func (e UnknownUserError) Error() string`

## type UnknownUserIdError

```
type UnknownUserIdError int
```

UnknownUserIdError is returned by LookupId when a user cannot be found.

当无法找到一个用户时，LookupId会返回UnknownUserIdError。

### func (UnknownUserIdError) Error

`func (e UnknownUserIdError) Error() string`

## type User

```
type User struct {
	// Uid is the user ID.
	// On POSIX systems, this is a decimal number representing the uid.
	// On Windows, this is a security identifier (SID) in a string format.
	// On Plan 9, this is the contents of /dev/user.
	Uid string
	// Gid is the primary group ID.
	// On POSIX systems, this is a decimal number representing the gid.
	// On Windows, this is a SID in a string format.
	// On Plan 9, this is the contents of /dev/user.
	Gid string
	// Username is the login name.
	Username string
	// Name is the user's real or display name.
	// It might be blank.
	// On POSIX systems, this is the first (or only) entry in the GECOS field
	// list.
	// On Windows, this is the user's display name.
	// On Plan 9, this is the contents of /dev/user.
	Name string
	// HomeDir is the path to the user's home directory (if they have one).
	HomeDir string
}
```
User represents a user account.

用户代表一个用户账户。

### func Current 

`func Current() (*User, error)`

Current returns the current user.

The first call will cache the current user information. Subsequent calls will return the cached value and will not reflect changes to the current user.

Current返回当前用户。

第一次调用将缓存当前用户信息。随后的调用将返回缓存的值，并且不反映对当前用户的改变。

### func Lookup

`func Lookup(username string) (*User, error)`

Lookup looks up a user by username. If the user cannot be found, the returned error is of type UnknownUserError.

Lookup通过用户名查找一个用户。如果不能找到该用户，返回的错误类型为未知用户错误。

### func LookupId

`func LookupId(uid string) (*User, error)`

LookupId looks up a user by userid. If the user cannot be found, the returned error is of type UnknownUserIdError.

LookupId通过userid查找一个用户。如果找不到该用户，返回的错误类型为UnknownUserIdError。

### func (*User) GroupIds 

`func (u *User) GroupIds() ([]string, error)`

GroupIds returns the list of group IDs that the user is a member of.

GroupIds返回用户是成员的群组ID的列表。