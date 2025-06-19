import unittest
from primo import es_primo

class TestPrimo(unittest.TestCase):
	def test_primos(self):
		primos = [2, 3, 5, 7, 11, 13, 17, 19, 23, 29]
		indice = 0
		longitud = len(primos)
		while indice < longitud:
			self.assertTrue(es_primo(primos[indice]))
			indice = indice + 1

	def test_no_primos(self):
		no_primos = [0, 1, 4, 6, 8, 9, 10, 12, 14, 15, 16, 18, 20]
		indice = 0
		longitud = len(no_primos)
		while indice < longitud:
			self.assertFalse(es_primo(no_primos[indice]))
			indice = indice + 1

	def test_extremos(self):
		negativos = [-10, -1, -100, -999]
		indice = 0
		longitud = len(negativos)
		while indice < longitud:
			self.assertFalse(es_primo(negativos[indice]))
			indice = indice + 1

		self.assertFalse(es_primo(0))
		self.assertFalse(es_primo(1))
		self.assertTrue(es_primo(7919))
		self.assertFalse(es_primo(7920))

if __name__ == "__main__":
	unittest.main() 