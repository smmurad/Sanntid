import threading

i = 0

lock = threading.Lock()

def threadFunction1():
	for k in range (1000002):
		lock.acquire()
		i += 1
		lock.release()

	global i

def threadFunction2():
	for k in range (1000000):
		lock.acquire()
		i-=1
		lock.release()
	global i


def main():
	
	#lock = threading.Lock()
	thread1 = threading.Thread(target = threadFunction1, args = (),)
	thread2 = threading.Thread(target = threadFunction2, args = (),)

	thread1.start()
	thread2.start()

	thread1.join()
	thread2.join()

	print("Value after both threads running i = ", i)

main()
