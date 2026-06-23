# Taller 3 - Algoritmo y Estructura de Datos (Universidad ESAN)

Este repositorio contiene la implementación en Go de diversas estructuras de datos clásicas aplicadas a problemas reales con procesamiento de grandes volúmenes de información.

 👥 Integrante(s)
Gael Almeyda Razo

 📹 Demostración y Pruebas de Ejecución
* 📁 **Evidencia de ejecución (Google Drive):https://drive.google.com/drive/folders/1nK2tMzqrm2lqkxFhuUt-z3Y6hlxUd9st
* 🔗 **Video Explicativo de la Defensa (YouTube):(https://youtu.be/LRjqdztENbw)
## 🏗️ Estructura del Repositorio
* `pilas/`: Cálculo de Stock Span utilizando un Stack Monótono en $O(n)$ (Dataset: `aapl.us.txt`).
* `colas/`: Implementación de un Rate Limiter para control de peticiones IP mediante una Cola FIFO (Dataset: `access.log`).
* `listas/`: Simulación de una Caché LRU de alta velocidad gestionada con una Lista Doblemente Enlazada y un Mapa.
* `arboles/`: Índice de búsqueda en Árbol AVL auto-balanceado para consultas por rango rápido (Dataset: MovieLens 100K).

## 🚀 Ejecución de Pruebas Unitarias (Testing)
El código incluye suites de pruebas unitarias (`*_test.go`) que garantizan el correcto comportamiento lógico y la estabilidad frente a casos límite (ej. balanceo autónomo del AVL, expulsión correcta en LRU). 

Para comprobar que todas las estructuras pasan sus pruebas al 100%, ejecutar el siguiente comando en la raíz del proyecto:
```bash
go test -v ./...

🤖 Declaración de Uso de IA
De acuerdo con los lineamientos del curso, se declara el uso de asistencia de IA (Gemini) como herramienta de apoyo metodológico para comprender la lógica matemática de las rotaciones (Izq/Der) en el Árbol AVL y para estructurar la plantilla base de los test unitarios. Todo el código fue validado empíricamente, depurado y comprobado mediante compilación exitosa.
