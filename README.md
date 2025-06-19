# Verificador de Números Primos - CLI

Una herramienta de línea de comandos para verificar si los números son primos.

## Características

- ✅ Verifica números primos usando algoritmo eficiente
- ✅ Soporte para múltiples números en una sola ejecución
- ✅ Manejo de errores para entradas inválidas
- ✅ Ayuda integrada
- ✅ Tests unitarios completos

## Uso

### Verificar un número:
```bash
python3 primo.py 17
# Salida: 17 es primo
```

### Verificar múltiples números:
```bash
python3 primo.py 17 18 19 23
# Salida: 
# 17 es primo
# 18 no es primo
# 19 es primo
# 23 es primo
```

### Mostrar ayuda:
```bash
python3 primo.py --help
```

## Tests

Ejecutar los tests unitarios:
```bash
python3 test_primo.py
```

## Implementación

- `es_primo(n)`: Función principal que determina si un número es primo
- Manejo robusto de argumentos CLI con `sys.argv`
- Validación de entrada y mensajes de error claros
- Tests exhaustivos incluyendo casos extremos

## Algoritmo

El algoritmo verifica divisibilidad hasta `n // 2` para optimizar el tiempo de ejecución, retornando `False` tan pronto como encuentra un divisor.