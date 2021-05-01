## <a name="how-do-i-use-go-mutesting"></a>Como usar go-mutesting

Instalar binário da biblioteca:

```bash
go get -t -v github.com/zimmski/go-mutesting/...
```

Abrir lista de ajuda:
```bash
go-mutesting --help
```

Os alvos do teste de mutação podem ser definidos como argumentos para o binário. Cada destino pode ser um arquivo de origem Go, um diretório ou um pacote. Os diretórios e pacotes também podem incluir o padrão curinga `...` que irá pesquisar recursivamente por arquivos fonte Go. Os arquivos de origem de teste com o sufixo `_test` são excluídos, pois isso interfere no processo de teste na maioria das vezes.

O exemplo a seguir reúne todos os arquivos Go que são definidos pelos destinos e geram mutações com todos os modificadores disponíveis do binário.
```bash
go-mutesting parse.go example/ github.com/zimmski/go-mutesting/mutator/...
```
