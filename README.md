# Ejemplo llamada a api externa

Llamo a la api de github para buscar repositorios y se guarda en el fichero csv que le digas.

Por ejemplo para hacer esta llamada:
```
https://api.github.com/search/repositories?q=airbnb/css
```

Se llama al cliente de la siguiente manera:
```
go run cmd/cli/main.go search -q=airbnb -o=/tmp/prueba.csv
```
