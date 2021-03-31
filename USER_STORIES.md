# Documento Lista de User Stories


## Descrição

Este documento descreve os User Stories criados a partir da Lista de Requisitos . 

## Histórico de revisões

| Data       | Versão  | Descrição                          | Autor                          |
| :--------- | :-----: | :--------------------------------: | :----------------------------- |
| 30/03/2021 | 0.0.1   | Descrição do documento  | Laio de Alencar |


### User Story US01 - Manter Usuário

|               |                                                                |
| ------------- | :------------------------------------------------------------- |
| **Descrição** | O sistema deve manter um cadastro de aluno que tem acesso ao sistema via login e senha. Um aluno tem os atributos name, id, username, data de nascimento, senha. O username será gerado ao cadastrar-lo no sistema. O usuário administrador do sistema pode realizar as operações de adicionar, alterar, remover e listar os usuários comuns do sistema. |

| **Requisitos envolvidos** |                                                    |
| ------------- | :------------------------------------------------------------- |
| RF01          | Cadastrar Aluno |
| RF02          | Alterar Aluno  |
| RF03          | Listar Aluno        |
| RF04          | Deletar Aluno |
| RF15          | Login/Logout |
|               |                                                                |
| ------------- | :------------------------------------------------------------- |
| **Prioridade**            | Essencial                           |
| **Estimativa**            |                                |
| **Tempo Gasto (real):**   |                                     |
| **Tamanho Funcional**     |                                 |

| Testes de Aceitação (TA) |  |
| ----------- | --------- |
| **Código**      | **Descrição** |
| **TA01.01** | 
O administrador informa, na tela Cadastrar, todos os dados para registrar-se corretamente, ao clicar em salvar ele é notificado com uma mensagem de sucesso. Mensagem: Cadastro realizado com sucesso.

 |
| **TA01.02** | 
O administrador informa, na tela Registrar, os dados para registrar-se incorretamente, ao clicar em salvar ele é notificado com uma mensagem de erro. Mensagem: Cadastro não realizado, o campo “xxxx” não foi informado corretamente.

 |
| **TA01.03** | 
O aluno informa, na tela Login, os dados para logar corretamente, ao clicar em entrar ele é direcionado para página inicial.

 |
| **TA01.04** | 
O aluno  informa, na tela Login, os dados para logar incorretamente, ao clicar em entrar ele é notificado com uma mensagem de erro.

 |