package ex5

func replaceSpace(str []byte, length int) {
	count := 0
	for i := 0; i < length; i++ {
		if str[i] == ' ' {
			count++
		}
	}
	//计算总的长度
	newlength := length + count*2

	for l, nl := length-1, newlength-1; l >= 0 && nl >= 0; {
		if str[l] == ' ' {
			str[nl] = '0'
			nl--
			str[nl] = '2'
			nl--
			str[nl] = '%'
			nl--
			l--
		} else {
			str[nl] = str[l]
			nl--
			l--
		}
	}
}
