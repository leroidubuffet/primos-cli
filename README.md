# Verificador de Números Primos - CLI

Una herramienta de línea de comandos para verificar si los números son primos, implementada en Go y Python.

## Características

- ✅ Verifica números primos usando algoritmo eficiente
- ✅ Soporte para múltiples números en una sola ejecución
- ✅ Manejo de errores para entradas inválidas
- ✅ Ayuda integrada
- ✅ Tests unitarios completos
- ✅ Implementaciones en Go y Python

## Uso

### Versión Go (Recomendada)

#### Verificar un número:
```bash
go run main.go 17
# Salida: 17 es primo
```

#### Verificar múltiples números:
```bash
go run main.go 17 18 19 23
# Salida: 
# 17 es primo
# 18 no es primo
# 19 es primo
# 23 es primo
```

#### Mostrar ayuda:
```bash
go run main.go --help
```

#### Compilar ejecutable:
```bash
go build -o primos main.go
./primos 17
```

### Versión Python

#### Verificar un número:
```bash
python3 primo.py 17
# Salida: 17 es primo
```

#### Verificar múltiples números:
```bash
python3 primo.py 17 18 19 23
```

#### Mostrar ayuda:
```bash
python3 primo.py --help
```

## Tests

### Tests Go:
```bash
go test -v                    # Tests con output detallado
go test -bench=.             # Benchmarks de rendimiento
```

### Tests Python:
```bash
python3 test_primo.py
```

## Implementación

### Go (main.go)
- `isPrime(n int) bool`: Función principal que determina si un número es primo
- Manejo robusto de argumentos CLI con `os.Args`
- Validación de entrada con `strconv.Atoi`
- Tests exhaustivos incluyendo benchmarks

### Python (primo.py)
- `es_primo(n)`: Función principal que determina si un número es primo
- Manejo robusto de argumentos CLI con `sys.argv`
- Validación de entrada y mensajes de error claros
- Tests exhaustivos incluyendo casos extremos

## Algoritmo

El algoritmo verifica divisibilidad hasta `n / 2` para optimizar el tiempo de ejecución, retornando `false` tan pronto como encuentra un divisor.

## Rendimiento

La versión en Go ofrece mejor rendimiento y tiempo de ejecución más rápido comparado con Python, especialmente para números grandes.