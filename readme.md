<h2>Command Line App</h2>

<h3>Este projeto foi criado durante os meus estudos sobre a linguagem Golang para fixação dos seus fundamentos, consiste em uma aplicação de linhas de comando.</h3>

<h3>A aplicação possui duas funcionalidades, buscar os ips públicos de um endereço e buscar os servidores onde estão hospedados.</h3>

<h4>Para executar a busca por ips da aplicação basta executar o comando:</h4>

    go run main.go ips --host nome_do_host

    
<h4>Exemplo:</h4>

    go run main.go ips --host amazon.com.br

<h4>O retorno será algo parecido com isto:</h4>

    52.94.225.243
    72.21.203.171
    54.239.26.87

<h4>Para executar a busca por servidores da aplicação basta executar o comando:</h4>

    go run main.go servers --host nome_do_host

<h4>Exemplo:</h4>

    go run main.go servers --host amazon.com.br

<h4>O retorno será algo parecido com isto:</h4>

    pdns5.ultradns.info.
    pdns4.ultradns.org.
    ns1.p31.dynect.net.
    pdns1.ultradns.net.
    ns2.p31.dynect.net.
    pdns3.ultradns.org.
