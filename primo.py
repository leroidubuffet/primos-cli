import sys

def es_primo(n):
	resultado = True
	contador = 2
	limite = n // 2
	if n < 2:
		resultado = False
	else:
		while contador <= limite and resultado:
			residuo = n % contador
			if residuo == 0:
				resultado = False
			else:
				contador = contador + 1
	return resultado

def mostrar_ayuda():
	print("Uso: python primo.py <numero1> [numero2] [numero3] ...")
	print("Ejemplo: python primo.py 17 23 100")
	print("Verifica si los números dados son primos")

def main():
	if len(sys.argv) < 2:
		print("Error: Se requiere al menos un número como argumento")
		mostrar_ayuda()
		sys.exit(1)
	
	if sys.argv[1] in ['-h', '--help', 'help']:
		mostrar_ayuda()
		sys.exit(0)
	
	for i in range(1, len(sys.argv)):
		try:
			numero_int = int(sys.argv[i])
			primo = es_primo(numero_int)
			if primo:
				print(f"{numero_int} es primo")
			else:
				print(f"{numero_int} no es primo")
		except ValueError:
			print(f"Error: '{sys.argv[i]}' no es un número válido")
			sys.exit(1)

if __name__ == "__main__":
	main() 