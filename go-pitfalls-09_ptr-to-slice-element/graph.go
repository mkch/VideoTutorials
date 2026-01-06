package main

/*
Before sorting:

	Members:

			｜ {3 Bob 30}     |    addr: 0xabc00
			｜ {2 Alice 25}   |    addr: 0xabc10
			｜ {1 Charlie 35} |    addr: 0xabc20

	ID Map:

			｜ 2: 0xabc10    |
			｜    ...        |
			｜ 3: 0xabc00    |
			｜    ...        |
			｜ 1: 0xabc20    |

After sorting by Name:

	Members:

			｜ {2 Alice 25}   |    addr: 0xabc00
			｜ {3 Bob 30}     |    addr: 0xabc10
			｜ {1 Charlie 35} |    addr: 0xabc20

	ID Map:
			｜ 2: 0xabc10    |
			｜    ...        |
			｜ 3: 0xabc00    |
			｜    ...        |
			｜ 1: 0xabc20    |
*/
