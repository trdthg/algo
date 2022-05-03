	.text
	.file	"1.a4659c16-cgu.0"
	.section	".text._ZN4core3ops8function6FnOnce40call_once$u7b$$u7b$vtable.shim$u7d$$u7d$17h54713602a7690b2cE","ax",@progbits
	.p2align	4, 0x90
	.type	_ZN4core3ops8function6FnOnce40call_once$u7b$$u7b$vtable.shim$u7d$$u7d$17h54713602a7690b2cE,@function
_ZN4core3ops8function6FnOnce40call_once$u7b$$u7b$vtable.shim$u7d$$u7d$17h54713602a7690b2cE:
	.cfi_startproc
	pushq	%rax
	.cfi_def_cfa_offset 16
	movq	(%rdi), %rdi
	callq	_ZN3std10sys_common9backtrace28__rust_begin_short_backtrace17h537af9aa8b92e5c8E
	xorl	%eax, %eax
	popq	%rcx
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end0:
	.size	_ZN4core3ops8function6FnOnce40call_once$u7b$$u7b$vtable.shim$u7d$$u7d$17h54713602a7690b2cE, .Lfunc_end0-_ZN4core3ops8function6FnOnce40call_once$u7b$$u7b$vtable.shim$u7d$$u7d$17h54713602a7690b2cE
	.cfi_endproc

	.section	".text._ZN4core3ptr85drop_in_place$LT$std..rt..lang_start$LT$$LP$$RP$$GT$..$u7b$$u7b$closure$u7d$$u7d$$GT$17h84d7731a2e75e375E","ax",@progbits
	.p2align	4, 0x90
	.type	_ZN4core3ptr85drop_in_place$LT$std..rt..lang_start$LT$$LP$$RP$$GT$..$u7b$$u7b$closure$u7d$$u7d$$GT$17h84d7731a2e75e375E,@function
_ZN4core3ptr85drop_in_place$LT$std..rt..lang_start$LT$$LP$$RP$$GT$..$u7b$$u7b$closure$u7d$$u7d$$GT$17h84d7731a2e75e375E:
	.cfi_startproc
	retq
.Lfunc_end1:
	.size	_ZN4core3ptr85drop_in_place$LT$std..rt..lang_start$LT$$LP$$RP$$GT$..$u7b$$u7b$closure$u7d$$u7d$$GT$17h84d7731a2e75e375E, .Lfunc_end1-_ZN4core3ptr85drop_in_place$LT$std..rt..lang_start$LT$$LP$$RP$$GT$..$u7b$$u7b$closure$u7d$$u7d$$GT$17h84d7731a2e75e375E
	.cfi_endproc

	.section	.text._ZN2_14main17h505d8a0902b0c21dE,"ax",@progbits
	.p2align	4, 0x90
	.type	_ZN2_14main17h505d8a0902b0c21dE,@function
_ZN2_14main17h505d8a0902b0c21dE:
	.cfi_startproc
	subq	$88, %rsp
	.cfi_def_cfa_offset 96
	leaq	.L__unnamed_1(%rip), %rax
	movq	%rax, 8(%rsp)
	movq	$5, 16(%rsp)
	leaq	8(%rsp), %rax
	movq	%rax, 24(%rsp)
	leaq	_ZN44_$LT$$RF$T$u20$as$u20$core..fmt..Display$GT$3fmt17h687204a082b63270E(%rip), %rax
	movq	%rax, 32(%rsp)
	leaq	.L__unnamed_2(%rip), %rax
	movq	%rax, 40(%rsp)
	movq	$2, 48(%rsp)
	movq	$0, 56(%rsp)
	leaq	24(%rsp), %rax
	movq	%rax, 72(%rsp)
	movq	$1, 80(%rsp)
	leaq	40(%rsp), %rdi
	callq	*_ZN3std2io5stdio6_print17hcbc8e5359e4501b6E@GOTPCREL(%rip)
	addq	$88, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end2:
	.size	_ZN2_14main17h505d8a0902b0c21dE, .Lfunc_end2-_ZN2_14main17h505d8a0902b0c21dE
	.cfi_endproc

	.section	.text._ZN3std2rt10lang_start17hf2b94e64f2a9671dE,"ax",@progbits
	.hidden	_ZN3std2rt10lang_start17hf2b94e64f2a9671dE
	.globl	_ZN3std2rt10lang_start17hf2b94e64f2a9671dE
	.p2align	4, 0x90
	.type	_ZN3std2rt10lang_start17hf2b94e64f2a9671dE,@function
_ZN3std2rt10lang_start17hf2b94e64f2a9671dE:
	.cfi_startproc
	pushq	%rax
	.cfi_def_cfa_offset 16
	movq	%rdx, %rcx
	movq	%rsi, %rdx
	movq	%rdi, (%rsp)
	leaq	.L__unnamed_3(%rip), %rsi
	movq	%rsp, %rdi
	callq	*_ZN3std2rt19lang_start_internal17h2ba92edce36c035eE@GOTPCREL(%rip)
	popq	%rcx
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end3:
	.size	_ZN3std2rt10lang_start17hf2b94e64f2a9671dE, .Lfunc_end3-_ZN3std2rt10lang_start17hf2b94e64f2a9671dE
	.cfi_endproc

	.section	".text._ZN3std2rt10lang_start28_$u7b$$u7b$closure$u7d$$u7d$17h78fc33a42296254fE","ax",@progbits
	.p2align	4, 0x90
	.type	_ZN3std2rt10lang_start28_$u7b$$u7b$closure$u7d$$u7d$17h78fc33a42296254fE,@function
_ZN3std2rt10lang_start28_$u7b$$u7b$closure$u7d$$u7d$17h78fc33a42296254fE:
	.cfi_startproc
	pushq	%rax
	.cfi_def_cfa_offset 16
	movq	(%rdi), %rdi
	callq	_ZN3std10sys_common9backtrace28__rust_begin_short_backtrace17h537af9aa8b92e5c8E
	xorl	%eax, %eax
	popq	%rcx
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end4:
	.size	_ZN3std2rt10lang_start28_$u7b$$u7b$closure$u7d$$u7d$17h78fc33a42296254fE, .Lfunc_end4-_ZN3std2rt10lang_start28_$u7b$$u7b$closure$u7d$$u7d$17h78fc33a42296254fE
	.cfi_endproc

	.section	.text._ZN3std10sys_common9backtrace28__rust_begin_short_backtrace17h537af9aa8b92e5c8E,"ax",@progbits
	.p2align	4, 0x90
	.type	_ZN3std10sys_common9backtrace28__rust_begin_short_backtrace17h537af9aa8b92e5c8E,@function
_ZN3std10sys_common9backtrace28__rust_begin_short_backtrace17h537af9aa8b92e5c8E:
	.cfi_startproc
	pushq	%rax
	.cfi_def_cfa_offset 16
	callq	*%rdi
	#APP
	#NO_APP
	popq	%rax
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end5:
	.size	_ZN3std10sys_common9backtrace28__rust_begin_short_backtrace17h537af9aa8b92e5c8E, .Lfunc_end5-_ZN3std10sys_common9backtrace28__rust_begin_short_backtrace17h537af9aa8b92e5c8E
	.cfi_endproc

	.section	".text._ZN44_$LT$$RF$T$u20$as$u20$core..fmt..Display$GT$3fmt17h687204a082b63270E","ax",@progbits
	.p2align	4, 0x90
	.type	_ZN44_$LT$$RF$T$u20$as$u20$core..fmt..Display$GT$3fmt17h687204a082b63270E,@function
_ZN44_$LT$$RF$T$u20$as$u20$core..fmt..Display$GT$3fmt17h687204a082b63270E:
	.cfi_startproc
	movq	%rsi, %rdx
	movq	(%rdi), %rax
	movq	8(%rdi), %rsi
	movq	%rax, %rdi
	jmpq	*_ZN42_$LT$str$u20$as$u20$core..fmt..Display$GT$3fmt17h4162dc8d40f3e60cE@GOTPCREL(%rip)
.Lfunc_end6:
	.size	_ZN44_$LT$$RF$T$u20$as$u20$core..fmt..Display$GT$3fmt17h687204a082b63270E, .Lfunc_end6-_ZN44_$LT$$RF$T$u20$as$u20$core..fmt..Display$GT$3fmt17h687204a082b63270E
	.cfi_endproc

	.section	.text.main,"ax",@progbits
	.globl	main
	.p2align	4, 0x90
	.type	main,@function
main:
	.cfi_startproc
	pushq	%rax
	.cfi_def_cfa_offset 16
	movq	%rsi, %rcx
	movslq	%edi, %rdx
	leaq	_ZN2_14main17h505d8a0902b0c21dE(%rip), %rax
	movq	%rax, (%rsp)
	leaq	.L__unnamed_3(%rip), %rsi
	movq	%rsp, %rdi
	callq	*_ZN3std2rt19lang_start_internal17h2ba92edce36c035eE@GOTPCREL(%rip)
	popq	%rcx
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end7:
	.size	main, .Lfunc_end7-main
	.cfi_endproc

	.type	.L__unnamed_1,@object
	.section	.rodata..L__unnamed_1,"a",@progbits
.L__unnamed_1:
	.zero	5,49
	.size	.L__unnamed_1, 5

	.type	.L__unnamed_4,@object
	.section	.rodata..L__unnamed_4,"a",@progbits
	.p2align	3
.L__unnamed_4:
	.size	.L__unnamed_4, 0

	.type	.L__unnamed_5,@object
	.section	.rodata..L__unnamed_5,"a",@progbits
.L__unnamed_5:
	.byte	10
	.size	.L__unnamed_5, 1

	.type	.L__unnamed_2,@object
	.section	.data.rel.ro..L__unnamed_2,"aw",@progbits
	.p2align	3
.L__unnamed_2:
	.quad	.L__unnamed_4
	.zero	8
	.quad	.L__unnamed_5
	.asciz	"\001\000\000\000\000\000\000"
	.size	.L__unnamed_2, 32

	.type	.L__unnamed_3,@object
	.section	.data.rel.ro..L__unnamed_3,"aw",@progbits
	.p2align	3
.L__unnamed_3:
	.quad	_ZN4core3ptr85drop_in_place$LT$std..rt..lang_start$LT$$LP$$RP$$GT$..$u7b$$u7b$closure$u7d$$u7d$$GT$17h84d7731a2e75e375E
	.asciz	"\b\000\000\000\000\000\000\000\b\000\000\000\000\000\000"
	.quad	_ZN4core3ops8function6FnOnce40call_once$u7b$$u7b$vtable.shim$u7d$$u7d$17h54713602a7690b2cE
	.quad	_ZN3std2rt10lang_start28_$u7b$$u7b$closure$u7d$$u7d$17h78fc33a42296254fE
	.quad	_ZN3std2rt10lang_start28_$u7b$$u7b$closure$u7d$$u7d$17h78fc33a42296254fE
	.size	.L__unnamed_3, 48

	.section	".note.GNU-stack","",@progbits
