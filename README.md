# nolint analyzer issue

This command should not raise any issue:

```console
$ go run main.go ./test
runtime: goroutine stack exceeds 1000000000-byte limit
runtime: sp=0xc0219023a8 stack=[0xc021902000, 0xc041902000]
fatal error: stack overflow

runtime stack:
runtime.throw({0x75a623?, 0x2030020?})
        /usr/lib/go/src/runtime/panic.go:1077 +0x5c fp=0xc000073e18 sp=0xc000073de8 pc=0x
43631c
runtime.newstack()
        /usr/lib/go/src/runtime/stack.go:1107 +0x5ac fp=0xc000073fc8 sp=0xc000073e18 pc=0
x45038c
traceback: unexpected SPWRITE function runtime.morestack
runtime.morestack()
        /usr/lib/go/src/runtime/asm_amd64.s:593 +0x8f fp=0xc000073fd0 sp=0xc000073fc8 pc=
0x46858f

goroutine 111 [running]:
runtime.heapBitsSetType(0xc0088dcea0, 0xd0, 0xd0, 0x74b920)
        /usr/lib/go/src/runtime/mbitmap.go:946 +0x576 fp=0xc0219023b8 sp=0xc0219023b0 pc=
0x414ef6
runtime.mallocgc(0xd0, 0x74b920, 0x1)
        /usr/lib/go/src/runtime/malloc.go:1144 +0x5bf fp=0xc021902420 sp=0xc0219023b8 pc=
0x40cdff
runtime.newobject(0xc000080000?)
        /usr/lib/go/src/runtime/malloc.go:1324 +0x25 fp=0xc021902448 sp=0xc021902420 pc=0
x40d1c5
main.main.WrapFunc.func2(0xc0088dcdd0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:44 +0x3e fp=0xc02
1902498 sp=0xc021902448 pc=0x6d1a3e
main.main.WrapFunc.func2(0xc0088dcd00)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
219024e8 sp=0xc021902498 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dcc30)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902538 sp=0xc0219024e8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dcb60)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902588 sp=0xc021902538 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dca90)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
219025d8 sp=0xc021902588 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dc9c0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902628 sp=0xc0219025d8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dc8f0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902678 sp=0xc021902628 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dc820)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
219026c8 sp=0xc021902678 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dc750)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902718 sp=0xc0219026c8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dc680)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902768 sp=0xc021902718 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dc5b0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
219027b8 sp=0xc021902768 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dc4e0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902808 sp=0xc0219027b8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dc410)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902858 sp=0xc021902808 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dc340)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
219028a8 sp=0xc021902858 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dc270)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
219028f8 sp=0xc0219028a8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dc1a0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902948 sp=0xc0219028f8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dc0d0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902998 sp=0xc021902948 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dc000)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
219029e8 sp=0xc021902998 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dbee0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902a38 sp=0xc0219029e8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dbe10)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902a88 sp=0xc021902a38 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dbd40)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902ad8 sp=0xc021902a88 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dbc70)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902b28 sp=0xc021902ad8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dbba0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902b78 sp=0xc021902b28 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dbad0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902bc8 sp=0xc021902b78 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dba00)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902c18 sp=0xc021902bc8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088db930)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902c68 sp=0xc021902c18 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088db860)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902cb8 sp=0xc021902c68 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088db790)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902d08 sp=0xc021902cb8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088db6c0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902d58 sp=0xc021902d08 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088db5f0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902da8 sp=0xc021902d58 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088db520)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902df8 sp=0xc021902da8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088db450)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902e48 sp=0xc021902df8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088db380)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902e98 sp=0xc021902e48 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088db2b0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902ee8 sp=0xc021902e98 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088db1e0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902f38 sp=0xc021902ee8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088db110)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902f88 sp=0xc021902f38 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088db040)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21902fd8 sp=0xc021902f88 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088daf70)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21903028 sp=0xc021902fd8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088daea0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21903078 sp=0xc021903028 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dadd0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
219030c8 sp=0xc021903078 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dad00)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21903118 sp=0xc0219030c8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dac30)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21903168 sp=0xc021903118 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088dab60)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
219031b8 sp=0xc021903168 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088daa90)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21903208 sp=0xc0219031b8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088da9c0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
21903258 sp=0xc021903208 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088da8f0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
219032a8 sp=0xc021903258 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0088da820)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
219032f8 sp=0xc0219032a8 pc=0x6d1b34
...6710773 frames elided...
main.main.WrapFunc.func2(0xc0003ccd00)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41900fd8 sp=0xc041900f88 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0003ccc30)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901028 sp=0xc041900fd8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0003ccb60)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901078 sp=0xc041901028 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0003cca90)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
419010c8 sp=0xc041901078 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0003cc9c0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901118 sp=0xc0419010c8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0003cc8f0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901168 sp=0xc041901118 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0003cc820)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
419011b8 sp=0xc041901168 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0003cc750)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901208 sp=0xc0419011b8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0003cc680)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901258 sp=0xc041901208 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0003cc5b0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
419012a8 sp=0xc041901258 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0003cc4e0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
419012f8 sp=0xc0419012a8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0003cc410)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901348 sp=0xc0419012f8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0003cc340)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901398 sp=0xc041901348 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0003cc270)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
419013e8 sp=0xc041901398 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0003cc1a0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901438 sp=0xc0419013e8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0003cc0d0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901488 sp=0xc041901438 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0003cc000)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
419014d8 sp=0xc041901488 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000221ee0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901528 sp=0xc0419014d8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000221e10)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901578 sp=0xc041901528 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000221d40)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
419015c8 sp=0xc041901578 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000221c70)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901618 sp=0xc0419015c8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000221ba0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901668 sp=0xc041901618 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000221ad0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
419016b8 sp=0xc041901668 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000221a00)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901708 sp=0xc0419016b8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000221930)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901758 sp=0xc041901708 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000221860)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
419017a8 sp=0xc041901758 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000221790)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
419017f8 sp=0xc0419017a8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0002216c0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901848 sp=0xc0419017f8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0002215f0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901898 sp=0xc041901848 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000221520)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
419018e8 sp=0xc041901898 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000221450)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901938 sp=0xc0419018e8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000221380)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901988 sp=0xc041901938 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0002212b0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
419019d8 sp=0xc041901988 pc=0x6d1b34
main.main.WrapFunc.func2(0xc0002211e0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901a28 sp=0xc0419019d8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000221110)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901a78 sp=0xc041901a28 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000221040)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901ac8 sp=0xc041901a78 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000220f70)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901b18 sp=0xc041901ac8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000220ea0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901b68 sp=0xc041901b18 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000220dd0)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901bb8 sp=0xc041901b68 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000220d00)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901c08 sp=0xc041901bb8 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000220c30)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901c58 sp=0xc041901c08 pc=0x6d1b34
main.main.WrapFunc.func2(0xc000220b60)
        /home/mgnsk/go/pkg/mod/github.com/kyoh86/nolint@v0.0.1/wrap.go:52 +0x134 fp=0xc0
41901ca8 sp=0xc041901c58 pc=0x6d1b34
golang.org/x/tools/go/analysis/internal/checker.(*action).execOnce(0xc0003c61e0)
        /home/mgnsk/go/pkg/mod/golang.org/x/tools@v0.14.0/go/analysis/internal/checker/c
hecker.go:767 +0x9f7 fp=0xc041901f18 sp=0xc041901ca8 pc=0x6966f7
golang.org/x/tools/go/analysis/internal/checker.(*action).execOnce-fm()
        <autogenerated>:1 +0x25 fp=0xc041901f30 sp=0xc041901f18 pc=0x698e85
sync.(*Once).doSlow(0xc000210d00?, 0xc0003ab518?)
        /usr/lib/go/src/sync/once.go:74 +0xbf fp=0xc041901f90 sp=0xc041901f30 pc=0x47f7bf
sync.(*Once).Do(...)
        /usr/lib/go/src/sync/once.go:65
golang.org/x/tools/go/analysis/internal/checker.(*action).exec(...)
        /home/mgnsk/go/pkg/mod/golang.org/x/tools@v0.14.0/go/analysis/internal/checker/c
hecker.go:683
golang.org/x/tools/go/analysis/internal/checker.execAll.func1(0x0?)
        /home/mgnsk/go/pkg/mod/golang.org/x/tools@v0.14.0/go/analysis/internal/checker/c
hecker.go:671 +0x45 fp=0xc041901fc8 sp=0xc041901f90 pc=0x695cc5
golang.org/x/tools/go/analysis/internal/checker.execAll.func2()
        /home/mgnsk/go/pkg/mod/golang.org/x/tools@v0.14.0/go/analysis/internal/checker/c
hecker.go:677 +0x27 fp=0xc041901fe0 sp=0xc041901fc8 pc=0x695c47
runtime.goexit()
        /usr/lib/go/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc041901fe8 sp=0xc041901fe0 pc=
0x46a2a1
created by golang.org/x/tools/go/analysis/internal/checker.execAll in goroutine 1
        /home/mgnsk/go/pkg/mod/golang.org/x/tools@v0.14.0/go/analysis/internal/checker/c
hecker.go:677 +0x15a

goroutine 1 [semacquire]:
runtime.gopark(0x7034c0?, 0xc0002cade0?, 0xe0?, 0x6a?, 0xc00035cea0?)
        /usr/lib/go/src/runtime/proc.go:398 +0xce fp=0xc000585918 sp=0xc0005858f8 pc=0x43
914e
runtime.goparkunlock(...)
        /usr/lib/go/src/runtime/proc.go:404
runtime.semacquire1(0xc000210d98, 0x18?, 0x1, 0x0, 0xab?)
        /usr/lib/go/src/runtime/sema.go:160 +0x218 fp=0xc000585980 sp=0xc000585918 pc=0x4
4a398
sync.runtime_Semacquire(0xc0003aba40?)
        /usr/lib/go/src/runtime/sema.go:62 +0x25 fp=0xc0005859b8 sp=0xc000585980 pc=0x466
a45
sync.(*WaitGroup).Wait(0x0?)
        /usr/lib/go/src/sync/waitgroup.go:116 +0x48 fp=0xc0005859e0 sp=0xc0005859b8 pc=0x
480b48
golang.org/x/tools/go/analysis/internal/checker.execAll({0xc00022aa60, 0x1, 0x62af76?})
        /home/mgnsk/go/pkg/mod/golang.org/x/tools@v0.14.0/go/analysis/internal/checker/c
hecker.go:680 +0x165 fp=0xc000585a38 sp=0xc0005859e0 pc=0x695be5
golang.org/x/tools/go/analysis/internal/checker.analyze({0xc00022aa58, 0x1, 0x1?}, {0xc00
0048428, 0x1, 0xc000135d90?})
        /home/mgnsk/go/pkg/mod/golang.org/x/tools@v0.14.0/go/analysis/internal/checker/c
hecker.go:311 +0x1f8 fp=0xc000585c28 sp=0xc000585a38 pc=0x692978
golang.org/x/tools/go/analysis/internal/checker.Run({0xc000014070, 0x1, 0x1}, {0xc0000484
28?, 0x1, 0x1})
        /home/mgnsk/go/pkg/mod/golang.org/x/tools@v0.14.0/go/analysis/internal/checker/c
hecker.go:146 +0x830 fp=0xc000585e30 sp=0xc000585c28 pc=0x692030
golang.org/x/tools/go/analysis/multichecker.Main({0xc000048428, 0x1, 0x1})
        /home/mgnsk/go/pkg/mod/golang.org/x/tools@v0.14.0/go/analysis/multichecker/multi
checker.go:59 +0x32e fp=0xc000585ee0 sp=0xc000585e30 pc=0x6d180e
main.main()
        /home/mgnsk/workspaces/nolint-issue/main.go:17 +0x136 fp=0xc000585f40 sp=0xc0005
85ee0 pc=0x6d19d6
runtime.main()
        /usr/lib/go/src/runtime/proc.go:267 +0x2bb fp=0xc000585fe0 sp=0xc000585f40 pc=0x4
38cfb
runtime.goexit()
        /usr/lib/go/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc000585fe8 sp=0xc000585fe0 pc=
0x46a2a1

goroutine 2 [force gc (idle)]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
        /usr/lib/go/src/runtime/proc.go:398 +0xce fp=0xc000044fa8 sp=0xc000044f88 pc=0x43
914e
runtime.goparkunlock(...)
        /usr/lib/go/src/runtime/proc.go:404
runtime.forcegchelper()
        /usr/lib/go/src/runtime/proc.go:322 +0xb3 fp=0xc000044fe0 sp=0xc000044fa8 pc=0x43
8fd3
runtime.goexit()
        /usr/lib/go/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc000044fe8 sp=0xc000044fe0 pc=
0x46a2a1
created by runtime.init.6 in goroutine 1
        /usr/lib/go/src/runtime/proc.go:310 +0x1a

goroutine 3 [GC sweep wait]:
runtime.gopark(0x1?, 0x0?, 0x0?, 0x0?, 0x0?)
        /usr/lib/go/src/runtime/proc.go:398 +0xce fp=0xc000045778 sp=0xc000045758 pc=0x43
914e
runtime.goparkunlock(...)
        /usr/lib/go/src/runtime/proc.go:404
runtime.bgsweep(0x0?)
        /usr/lib/go/src/runtime/mgcsweep.go:321 +0xdf fp=0xc0000457c8 sp=0xc000045778 pc=
0x4246ff
runtime.gcenable.func1()
        /usr/lib/go/src/runtime/mgc.go:200 +0x25 fp=0xc0000457e0 sp=0xc0000457c8 pc=0x419
845
runtime.goexit()
        /usr/lib/go/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc0000457e8 sp=0xc0000457e0 pc=
0x46a2a1
created by runtime.gcenable in goroutine 1
        /usr/lib/go/src/runtime/mgc.go:200 +0x66

goroutine 4 [GC scavenge wait]:
runtime.gopark(0x1527ded?, 0x14bca27?, 0x0?, 0x0?, 0x0?)
        /usr/lib/go/src/runtime/proc.go:398 +0xce fp=0xc000045f70 sp=0xc000045f50 pc=0x43
914e
runtime.goparkunlock(...)
        /usr/lib/go/src/runtime/proc.go:404
runtime.(*scavengerState).park(0x988e80)
        /usr/lib/go/src/runtime/mgcscavenge.go:425 +0x49 fp=0xc000045fa0 sp=0xc000045f70
pc=0x421f49
runtime.bgscavenge(0x0?)
        /usr/lib/go/src/runtime/mgcscavenge.go:658 +0x59 fp=0xc000045fc8 sp=0xc000045fa0
pc=0x4224f9
runtime.gcenable.func2()
        /usr/lib/go/src/runtime/mgc.go:201 +0x25 fp=0xc000045fe0 sp=0xc000045fc8 pc=0x419
7e5
runtime.goexit()
        /usr/lib/go/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc000045fe8 sp=0xc000045fe0 pc=
0x46a2a1
created by runtime.gcenable in goroutine 1
        /usr/lib/go/src/runtime/mgc.go:201 +0xa5

goroutine 5 [finalizer wait]:
runtime.gopark(0x0?, 0x770990?, 0x0?, 0x60?, 0x2000000020?)
        /usr/lib/go/src/runtime/proc.go:398 +0xce fp=0xc000044628 sp=0xc000044608 pc=0x43
914e
runtime.runfinq()
        /usr/lib/go/src/runtime/mfinal.go:193 +0x107 fp=0xc0000447e0 sp=0xc000044628 pc=0
x4188c7
runtime.goexit()
        /usr/lib/go/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc0000447e8 sp=0xc0000447e0 pc=
0x46a2a1
created by runtime.createfing in goroutine 1
        /usr/lib/go/src/runtime/mfinal.go:163 +0x3d

goroutine 112 [GC worker (idle)]:
runtime.gopark(0x156881cd831d?, 0x1?, 0x23?, 0x16?, 0xc0003c61e0?)
        /usr/lib/go/src/runtime/proc.go:398 +0xce fp=0xc000284f50 sp=0xc000284f30 pc=0x43
914e
runtime.gcBgMarkWorker()
        /usr/lib/go/src/runtime/mgc.go:1293 +0xe5 fp=0xc000284fe0 sp=0xc000284f50 pc=0x41
b3c5
runtime.goexit()
        /usr/lib/go/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc000284fe8 sp=0xc000284fe0 pc=
0x46a2a1
created by runtime.gcBgMarkStartWorkers in goroutine 111
        /usr/lib/go/src/runtime/mgc.go:1217 +0x1c

goroutine 177 [GC worker (idle)]:
runtime.gopark(0x9b9100?, 0x1?, 0xa7?, 0x9d?, 0xc000066e40?)
        /usr/lib/go/src/runtime/proc.go:398 +0xce fp=0xc000284750 sp=0xc000284730 pc=0x43
914e
runtime.gcBgMarkWorker()
        /usr/lib/go/src/runtime/mgc.go:1293 +0xe5 fp=0xc0002847e0 sp=0xc000284750 pc=0x41
b3c5
runtime.goexit()
        /usr/lib/go/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc0002847e8 sp=0xc0002847e0 pc=
0x46a2a1
created by runtime.gcBgMarkStartWorkers in goroutine 111
        /usr/lib/go/src/runtime/mgc.go:1217 +0x1c

goroutine 166 [GC worker (idle)]:
runtime.gopark(0x156881cd55f5?, 0xc000361790?, 0xa?, 0x76?, 0x480a45?)
        /usr/lib/go/src/runtime/proc.go:398 +0xce fp=0xc000361750 sp=0xc000361730 pc=0x43
914e
runtime.gcBgMarkWorker()
        /usr/lib/go/src/runtime/mgc.go:1293 +0xe5 fp=0xc0003617e0 sp=0xc000361750 pc=0x41
b3c5
runtime.goexit()
        /usr/lib/go/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc0003617e8 sp=0xc0003617e0 pc=
0x46a2a1
created by runtime.gcBgMarkStartWorkers in goroutine 111
        /usr/lib/go/src/runtime/mgc.go:1217 +0x1c

goroutine 178 [GC worker (idle)]:
runtime.gopark(0x9b9100?, 0x1?, 0x8b?, 0x41?, 0x0?)
        /usr/lib/go/src/runtime/proc.go:398 +0xce fp=0xc000281750 sp=0xc000281730 pc=0x43
914e
runtime.gcBgMarkWorker()
        /usr/lib/go/src/runtime/mgc.go:1293 +0xe5 fp=0xc0002817e0 sp=0xc000281750 pc=0x41
b3c5
runtime.goexit()
        /usr/lib/go/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc0002817e8 sp=0xc0002817e0 pc=
0x46a2a1
created by runtime.gcBgMarkStartWorkers in goroutine 111
        /usr/lib/go/src/runtime/mgc.go:1217 +0x1c
exit status 2
```
