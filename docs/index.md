<p style="text-align:center">
    <b>nohi</b><br/>
    <em>Fast command-line tool and library for generation human ids, uuids and passwords.</em><br/><br/>
    <a href="https://gitlab.com/unitto/nohi/-/commits/master"><img alt="coverage report" src="https://gitlab.com/unitto/nohi/badges/master/coverage.svg" /></a>
    <a href="https://gitlab.com/unitto/nohi/-/commits/master"><img alt="pipeline status" src="https://gitlab.com/unitto/nohi/badges/master/pipeline.svg" /></a>
    <a href="https://pkg.go.dev/gitlab.com/unitto/nohi/"><img src="https://pkg.go.dev/badge/badge/gitlab.com/unitto/nohi/" alt="PkgGoDev"></a>
</p>

---
# Introduction
> **Last Release:** v0.0.0

> **Docs Commit:** {commit}

> **Docs URL:** [nohi.unittolabs.com](https://nohi.unittolabs.com)

_NOHI - it's an abbreviation from Not Only Human-friendly IDs_

The main idea of this library and command-lite tool is to provide easy and performant
way to generate useful sensitive data such a Human-friendly ID, password, etc.

## Requirements
### Build from the source
[go >= 1.14](https://golang.org/doc/install)

make >= 4.3 (`sudo [apt|dnf|eopkg] install make`, check your distro docs)

### Pre-built docker image
[docker >= 19.03](https://docs.docker.com/engine/install/)

## Installation
### Build from the source

```bash
# clone source repository
$ git clone git@gitlab.com:unitto/nohi.git

# go to cloned dir
$ cd nohi

# use project shortcut
$ make build

# change user permissions
$ sudo chowd root:root nohi

# copy binary to the path
$ sudo mv nohi /usr/bin/

# test instalation
$ nohi -h

# to install auto-completion check the
$ nohi completion -h
```

### Pre-built docker image

```bash
# run docker image
$ docker run unitto/nohi -h
```

## Example Usage
Lest generate 10 logins and passwords for our mates:

```bash
$ nohi hid -n 10 && nohi pwd -n 10
boring_swartz_6zbdj
charming_wright_w9acu
clever_sutherland_82o4j
stoic_pascal_isa3v
vibrant_black_agcyb
hopeful_jennings_3uhu8
elated_borg_cgi0h
ecstatic_swartz_2nv8f
goofy_ritchie_c2366
sharp_ardinghelli_4w3tm

Done
c66k77ojt8lu7makk4dc0e8pjrl7oukv
ihn36voltcjz89njk3omyxamr1cbrubx
sll86ixpwi9fvpldok42dj4kg7qkue6j
qd6193vcdft4hy9zufu3sqhs6a9jtsrg
pcakkelrxmt2ggnd0a2q4mfp2jf1wfxr
syig9481g4d3axb2fca2gnqiq22abbuj
xtnlv5iz3d9cd8465oqovdspv70dw2j4
x7liiwyc4hjnbfe0vmk9400l6kp2usmm
s1ie3a52fwhyr95871hh5l17u931uy72
l1sypc7z8fk4q50qg78z50sfh6b5mupb

Done
```

## Dependencies

* [github.com/google/uuid](https://github.com/google/uuid) v1.1.1
* [github.com/spf13/cobra](https://github.com/spf13/cobra) v1.0.0

## Performance
> **NOTE:** this is not the best performance but good enough for
> most situations (it will be improved in future)

### Hardware
```bash
$ lscpu
Architecture:        x86_64
CPU op-mode(s):      32-bit, 64-bit
Byte Order:          Little Endian
Address sizes:       39 bits physical, 48 bits virtual
CPU(s):              4
On-line CPU(s) list: 0-3
Thread(s) per core:  1
Core(s) per socket:  4
Socket(s):           1
NUMA node(s):        1
Vendor ID:           GenuineIntel
CPU family:          6
Model:               60
Model name:          Intel(R) Core(TM) i5-4670K CPU @ 3.40GHz
Stepping:            3
CPU MHz:             1609.482
CPU max MHz:         3800.0000
CPU min MHz:         800.0000
BogoMIPS:            6800.58
Virtualization:      VT-x
L1d cache:           32K
L1i cache:           32K
L2 cache:            256K
L3 cache:            6144K
NUMA node0 CPU(s):   0-3
Flags:               fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm constant_tsc arch_perfmon pebs bts rep_good nopl xtopology nonstop_tsc cpuid aperfmperf pni pclmulqdq dtes64 monitor ds_cpl vmx est tm2 ssse3 sdbg fma cx16 xtpr pdcm pcid sse4_1 sse4_2 movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand lahf_lm abm cpuid_fault invpcid_single pti ssbd ibrs ibpb stibp tpr_shadow vnmi flexpriority ept vpid ept_ad fsgsbase tsc_adjust bmi1 avx2 smep bmi2 erms invpcid xsaveopt dtherm ida arat pln pts md_clear flush_l1d
```

### Quick Benchmark
```bash
$ make bench
go test -v -benchmem ./... -bench .
?       gitlab.com/unitto/nohi  [no test files]
?       gitlab.com/unitto/nohi/cmd      [no test files]
?       gitlab.com/unitto/nohi/internal/utils   [no test files]
=== RUN   TestAdapter
--- PASS: TestAdapter (0.00s)
PASS
ok      gitlab.com/unitto/nohi/pkg/adapter      0.002s
=== RUN   TestHumanic
--- PASS: TestHumanic (0.00s)
goos: linux
goarch: amd64
pkg: gitlab.com/unitto/nohi/pkg/hid
BenchmarkLoop
BenchmarkLoop-4              310           3753568 ns/op          527365 B/op      10000 allocs/op
BenchmarkBulk
BenchmarkBulk-4              454           2513104 ns/op          543913 B/op      10002 allocs/op
BenchmarkWait
BenchmarkWait-4              484           2570427 ns/op          560241 B/op      10003 allocs/op
PASS
ok      gitlab.com/unitto/nohi/pkg/hid  4.460s
=== RUN   TestPWD
--- PASS: TestPWD (0.00s)
goos: linux
goarch: amd64
pkg: gitlab.com/unitto/nohi/pkg/pwd
BenchmarkLoop
BenchmarkLoop-4              549           2170573 ns/op          640000 B/op       5000 allocs/op
BenchmarkBulk
BenchmarkBulk-4              778           1577485 ns/op          656514 B/op       5002 allocs/op
BenchmarkWait
BenchmarkWait-4              702           1617472 ns/op          672880 B/op       5003 allocs/op
PASS
ok      gitlab.com/unitto/nohi/pkg/pwd  4.107s
=== RUN   TestUUID4
--- PASS: TestUUID4 (0.00s)
goos: linux
goarch: amd64
pkg: gitlab.com/unitto/nohi/pkg/uuid4
BenchmarkLoop
BenchmarkLoop-4             1363            884283 ns/op           64000 B/op       2000 allocs/op
BenchmarkBulk
BenchmarkBulk-4             1924            595256 ns/op           80499 B/op       2002 allocs/op
BenchmarkWait
BenchmarkWait-4             1980            611375 ns/op           96869 B/op       2003 allocs/op
PASS
ok      gitlab.com/unitto/nohi/pkg/uuid4        3.780s
```

### Long Benchmark
```bash
$ make bench-long 
go test -v -cpu 1,2,4,8 -benchmem -benchtime 5s ./... -bench .
?       gitlab.com/unitto/nohi  [no test files]
?       gitlab.com/unitto/nohi/cmd      [no test files]
?       gitlab.com/unitto/nohi/internal/utils   [no test files]
=== RUN   TestAdapter
--- PASS: TestAdapter (0.00s)
=== RUN   TestAdapter
--- PASS: TestAdapter (0.00s)
=== RUN   TestAdapter
--- PASS: TestAdapter (0.00s)
=== RUN   TestAdapter
--- PASS: TestAdapter (0.00s)
PASS
ok      gitlab.com/unitto/nohi/pkg/adapter      0.002s
=== RUN   TestHumanic
--- PASS: TestHumanic (0.00s)
=== RUN   TestHumanic
--- PASS: TestHumanic (0.00s)
=== RUN   TestHumanic
--- PASS: TestHumanic (0.00s)
=== RUN   TestHumanic
--- PASS: TestHumanic (0.00s)
goos: linux
goarch: amd64
pkg: gitlab.com/unitto/nohi/pkg/hid
BenchmarkLoop
BenchmarkLoop               1598           3621332 ns/op          527365 B/op      10000 allocs/op
BenchmarkLoop-2             1598           3677178 ns/op          527365 B/op      10000 allocs/op
BenchmarkLoop-4             1650           3619226 ns/op          527362 B/op      10000 allocs/op
BenchmarkLoop-8             1634           3637374 ns/op          527367 B/op      10000 allocs/op
BenchmarkBulk
BenchmarkBulk               1672           3583098 ns/op          543845 B/op      10002 allocs/op
BenchmarkBulk-2             2192           2764511 ns/op          543852 B/op      10002 allocs/op
BenchmarkBulk-4             2368           2506889 ns/op          543863 B/op      10002 allocs/op
BenchmarkBulk-8             2332           2551464 ns/op          543870 B/op      10002 allocs/op
BenchmarkWait
BenchmarkWait               1660           3597848 ns/op          560226 B/op      10003 allocs/op
BenchmarkWait-2             2194           2755630 ns/op          560231 B/op      10003 allocs/op
BenchmarkWait-4             2445           2509566 ns/op          560231 B/op      10003 allocs/op
BenchmarkWait-8             2335           2567725 ns/op          560240 B/op      10003 allocs/op
PASS
ok      gitlab.com/unitto/nohi/pkg/hid  75.480s
=== RUN   TestPWD
--- PASS: TestPWD (0.00s)
=== RUN   TestPWD
--- PASS: TestPWD (0.00s)
=== RUN   TestPWD
--- PASS: TestPWD (0.00s)
=== RUN   TestPWD
--- PASS: TestPWD (0.00s)
goos: linux
goarch: amd64
pkg: gitlab.com/unitto/nohi/pkg/pwd
BenchmarkLoop
BenchmarkLoop               2815           2128779 ns/op          640000 B/op       5000 allocs/op
BenchmarkLoop-2             2802           2147529 ns/op          640000 B/op       5000 allocs/op
BenchmarkLoop-4             2811           2122222 ns/op          640000 B/op       5000 allocs/op
BenchmarkLoop-8             2769           2137422 ns/op          640003 B/op       5000 allocs/op
BenchmarkBulk
BenchmarkBulk               2812           2115896 ns/op          656480 B/op       5002 allocs/op
BenchmarkBulk-2             3562           1685246 ns/op          656484 B/op       5002 allocs/op
BenchmarkBulk-4             3838           1562762 ns/op          656487 B/op       5002 allocs/op
BenchmarkBulk-8             3705           1601187 ns/op          656493 B/op       5002 allocs/op
BenchmarkWait
BenchmarkWait               2752           2158775 ns/op          672864 B/op       5003 allocs/op
BenchmarkWait-2             3542           1688532 ns/op          672864 B/op       5003 allocs/op
BenchmarkWait-4             3745           1579149 ns/op          672865 B/op       5003 allocs/op
BenchmarkWait-8             3742           1598908 ns/op          672869 B/op       5003 allocs/op
PASS
ok      gitlab.com/unitto/nohi/pkg/pwd  73.897s
=== RUN   TestUUID4
--- PASS: TestUUID4 (0.00s)
=== RUN   TestUUID4
--- PASS: TestUUID4 (0.00s)
=== RUN   TestUUID4
--- PASS: TestUUID4 (0.00s)
=== RUN   TestUUID4
--- PASS: TestUUID4 (0.00s)
goos: linux
goarch: amd64
pkg: gitlab.com/unitto/nohi/pkg/uuid4
BenchmarkLoop
BenchmarkLoop               6896            859356 ns/op           64000 B/op       2000 allocs/op
BenchmarkLoop-2             6500            878872 ns/op           64000 B/op       2000 allocs/op
BenchmarkLoop-4             6858            863327 ns/op           64000 B/op       2000 allocs/op
BenchmarkLoop-8             6890            864284 ns/op           64000 B/op       2000 allocs/op
BenchmarkBulk
BenchmarkBulk               6499            910183 ns/op           80480 B/op       2002 allocs/op
BenchmarkBulk-2             8583            720365 ns/op           80482 B/op       2002 allocs/op
BenchmarkBulk-4            10254            583933 ns/op           80481 B/op       2002 allocs/op
BenchmarkBulk-8             8658            624016 ns/op           80487 B/op       2002 allocs/op
BenchmarkWait
BenchmarkWait               6506            912405 ns/op           96864 B/op       2003 allocs/op
BenchmarkWait-2             8446            732361 ns/op           96865 B/op       2003 allocs/op
BenchmarkWait-4             8884            590885 ns/op           96866 B/op       2003 allocs/op
BenchmarkWait-8             9261            632138 ns/op           96866 B/op       2003 allocs/op
PASS
ok      gitlab.com/unitto/nohi/pkg/uuid4        75.442s
```

---
<p style="text-align:center">
    <em>nohi is [MIT](https://choosealicense.com/licenses/mit/) licensed code. Designed & built in Kiev, Ukraine.</em>
</p>
