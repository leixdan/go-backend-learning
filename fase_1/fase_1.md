# 🚀 Cápsula 1: El Monolito Aislado (CLI Calculator + Docker)
Vamos a fusionar la lógica básica con tu deseo de usar contenedores desde el principio. No haremos una simple suma; haremos un programa de terminal (CLI) robusto que se ejecute dentro de un contenedor.

🎯 Tu Misión:
Crear una calculadora por línea de comandos que reciba tres argumentos (número 1, operación, número 2), devuelva el resultado manejando errores, y que se ejecute a través de Docker.

🛠️ Requisitos Técnicos a implementar:

Inicializar el módulo: Crea una carpeta para el proyecto y ejecuta go mod init github.com/tu-usuario/cli-calc.

Entrada de datos: Usa el paquete os (específicamente os.Args) para leer lo que el usuario escribe en la terminal.

Conversión: Los argumentos de la terminal entran como texto (string). Usa el paquete strconv para convertirlos a números de punto flotante (float64). Atención aquí: strconv.ParseFloat devuelve un valor y un error. Debes manejar ese error.

Lógica de la Calculadora: Crea una función llamada calculate(a float64, operator string, b float64) (float64, error).

Debe soportar: +, -, *, /.

El toque Pro: Si el operador es / y b es 0, debes devolver un error explícito usando el paquete errors (ej. errors.New("division by zero is not allowed")).

Dockerización Nivel 1: Crea un archivo llamado Dockerfile (sin extensión) en la raíz de tu proyecto. Tu objetivo es escribir las instrucciones para que Docker:

Use una imagen base de Go (ej. golang:1.22-alpine).

Copie tu código dentro del contenedor.

Compile tu código generando un binario.

Defina ese binario como el punto de entrada (ENTRYPOINT).

💡 Ejemplo de cómo debería funcionar al final:

Si compilaras en local y lo corrieras:
go run main.go 10 + 5 -> Salida: Result: 15
go run main.go 10 / 0 -> Salida: Error: division by zero is not allowed

Pero tu objetivo final es correrlo con Docker. Deberías poder construir la imagen y correrla así:
docker build -t cli-calc .
docker run --rm cli-calc 10 * 5 -> Salida: Result: 50