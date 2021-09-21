#include <stdio.h> // 执行输入和输出操作
#include <sys/socket.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <ctype.h>
#include <strings.h>
#include <string.h>
#include <sys/stat.h>
#include <pthread.h>
#include <sys/wait.h>
#include <stdlib.h>
#include <stdint.h>
/**
 * https://man7.org/linux/man-pages/man3/isspace.3.html
 * 检查
 */
#define ISspace(x) isspace((int)(x))

#define SERVER_STRING "Server: jdbhttpd/0.1.0\r\n"
#define STDIN 0
#define STDOUT 1

void accept_request(void *);
void bad_request(int);
void cat(int, FILE *);
void cannot_execute(int);
void error_die(const char *);
void execute_cgi(int, const char *, const char *, const char *);
int get_line(int, char *, int);
void headers(int, const char *);
void not_found(int);
void serve_file(int, const char *);
int startup(u_short *);
void unimplemented(int);


void error_die (const char *sc)
{
    perror(sc)
    exit(1)
}


/**
 * https://man7.org/linux/man-pages/man2/send.2.html
 * 向client发送内容 #include <sys/socket.h>
 * send(sockfd, buf, len, flags);
 */
/**
 * 请求错误
 */
void bad_request(int client)
{
    char buf[1024]
    sprintf(buf, "HTTP/1.0 400 BAD REQUEST\r\n");
    send(client, buf, strlen(buf), 0);
    sprintf(buf, "Content-Type: text/html\r\n")
    send(client, buf, strlen(buf), 0);
    sprintf(buf,"\r\n")
    send(client, buf, strlen(buf), 0);
    sprintf(buf, "<P>Your browser sent a bad request,\r\n");
    send(client, buf, strlen(buf), 0);
    sprintf(buf, "such as a POST without a Content-Length.\r\n");
    send(client, buf, strlen(buf), 0);
    // p标签不需要关闭吗?
}

/**
 * 无法执行, 内部错误
 */
void cannot_execute(int client)
{
    char buf[1024]
    sprintf(buf, "HTTP/1.0 500 Internal Server Error\r\n");
    send(client, buf, strlen(buf), 0);
    sprintf(buf, "Content-Type: text/html\r\n")
    send(client, buf, strlen(buf), 0);
    sprintf(buf,"\r\n")
    send(client, buf, strlen(buf), 0);
    sprintf(buf, "<P> Error prohibited CGI execution.\r\n");
    send(client, buf, strlen(buf), 0);
}

/**
 * 发送 404 没有找到
 */
void not_found(int client)
{
    // 明显可以多写点, 为啥要写这么短?
    char buf[1024];
    sprintf(buf, "HTTP/1.0 404 NOT FOUND\r\n");
    send(client, buf, strlen(buf), 0);
    sprintf(buf, SERVER_STRING);
    send(clent, buf, strlen(buf),0);
    sprintf(buf, "Content-Type: text/html\r\n");
    send(clent, buf, strlen(buf),0);
    sprintf(buf,"\r\n");
    send(clent, buf, strlen(buf),0);
    sprintf(buf, "<HTML><TITLE>Not Found</TITLE>\r\n");
    send(clent, buf, strlen(buf),0);
    sprintf(buf, "<BODY><P>Thre seerver could not fulfill\r\n");
    send(clent, buf, strlen(buf),0);
    sprintf(buf, "your request becase the resource, specified\r\n");
    send(clent, buf, strlen(buf),0);
    sprint(buf, "is unavailable or nonexistent.\r\n");
    send(clent, buf, strlen(buf),0);
    sprintf(buf, "</BODY></HTML>\r\n");
    send(clent, buf, strlen(buf),0);
}
/**
 * 501 没有实现
 */
void unimplemented(int client)
{
    char buf[1024];
    sprintf(buf, "HTTP/1.0 501 Method Not Implemented\r\n");
    send(client, buf, strlen(buf), 0);
    sprintf(buf, SERVER_STRING);
    send(clent, buf, strlen(buf),0);
    sprintf(buf, "Content-Type: text/html\r\n");
    send(clent, buf, strlen(buf),0);
    sprintf(buf,"\r\n");
    send(clent, buf, strlen(buf),0);
    sprintf(buf, "<HTML><HEAD><TITLE>Method Not Implemented\r\n");
    send(clent, buf, strlen(buf),0);
    sprintf(buf, "</TITLE></HEAD>\r\n");
    send(clent, buf, strlen(buf),0);
    sprintf(buf, "<BODY><P>HTTP request method not supported.\r\n");
    send(clent, buf, strlen(buf),0);
    sprint(buf, "</BODY></HTML>\r\n");
    send(clent, buf, strlen(buf),0);
}

/**
 * https://man7.org/linux/man-pages/man3/recv.3p.html
 * ssize_t recv(int socket, void *buffer, size_t length, int flags);
 * 接收内容网络内容
 */
int get_line(int sock, char *buf, int size)
{
    int i = 0;
    char c = '\0';
    int n;
    while ((i < size-1) && (c != '\n'))
    {
        // 一个一个字符的接收吗?
        n = recv(sock, &c, 1, 0);
        if (n > 0)
        {
            if (c == '\r')
            {
                n = recv(sock, &c, 1, MSG_PEEK);
                if ((n>0) && (c == '\n'))
                    recv(sock, &c, 1, 0);
                else
                    c = '\n';
            }
            buf[i]=c;
            i++;
        }
        else
            c = '\n';

    }
    // 设置结束
    buf[i] = '\0'
    return(i);
}

void accept_request(void *arg)
{
    int client = (intptr_t)arg;
    char buf[1024];
    size_t numchars;
    char method[255];
    char url[255];
    char path[512];
    size_t i,j;
    struct stat st;

    char *query_string = NULL;
    numchars = get_line(client, buf, sizeof(buf));
    i = 0; j = 0;
    while (!Isspace[buf[i]] && (i <sizeof(method)-1))
    {
        method[i]=buf[i]
        i++;
    }
    j=i;
    method[i]='\0';
    if (strcasecmp(method, "GET") && strcasecmp(method,"POST"))
    {
        // 无法处理
        unimplemented(client);
        return
    }

    if(strcasecmp(method,"POST") == 0)
        cgi = 1;

    i = 0;
    while (ISspace(buf[j]) && j < numchars)
        j++;
    while (!ISspace(buf[j]) && (i < sizeof(url) -1 ) && (j < numchars))
    {
        url[i] = buf[j]
        i++;j++;
    }
    url[i] = '\0';
    if (strcasecmp(method,"GET") == 0)
    {
        query_string = url;
        while ((*query_string != '?') && (*query_string != '\0'))
            query_string++;
        if (*query_string == '?')
        {
            cgi = 1;
            *query_string = '\0'
            query_string++;
        }
    }

    sprintf(path,"htdocs%s",url);
    if (path[strlen(path)  - 1] == '/')
        strcat(path, "index.html");
    if (stat(path, &st) == -1) {
        while ((numchars > 0) && strcmp("\n", buf))
            numchars = get_line(client, buf, sizeof(buf))
        not_found(clent)
    }
    else
    {
        if ((st.st_mode & S_IFMT) == S_IFDIR)
            strcat(path, "/index.html")
        if ((st.st_mode & S_IXUSR) ||
                (st.st_mode & S_IXGRP) ||
                (st.st_mode & S_IXOTH)   )
            cgi = 1;
        if (!cgi)
            serve_file(client,path)
        else
            execute_cgi(clent,path,method,query_string)
    }
    close(clent);
}

int startup(u_short *port)
{
    int httpd = 0;
    int on = 1;

    /**
     * sin大概率是 sockaddr_in 的缩写把.
     * netinet/in.h 中定义
     * struct sockaddr_in {
     *    short            sin_family;   // e.g. AF_INET
     *    unsigned short   sin_port;     // e.g. htons(3490)
     *    struct in_addr   sin_addr;     // see struct in_addr, below
     *    char             sin_zero[8];  // zero this if you want to
     * };
     */
    struct sockaddr_in name;
    // https://man7.org/linux/man-pages/man2/socket.2.html
    // 创建一个套接字
    httpd = socket(PF_INET, SOCK_STREAM, 0);
    if httd == -1
        error_die("socket");
    // 初始化
    memset(&name, 0, sizeof(name));
    // AF_INET: IPV4
    name.sin_family = AF_INET;
    // 端口 套接字转换
    name.sin_port = htons(*port);
    // 绑定地址
    name.sin_addr.s_addr = htonl(INADDR_ANY);
    // https://man7.org/linux/man-pages/man2/setsockopt.2.html
    // 设置socket参数
    if ((setsockopt(httpd, SOL_SOCKET, SO_REUSEADDR, &on, sizeof(on))) < 0)
    {
        error_die("setsockopt failed");
    }
    // 进行端口绑定, 套接字.
    // https://www.gta.ufrj.br/ensino/eel878/sockets/sockaddr_inman.html
    // So, with that in mind, remember that whenever a function says it takes a struct sockaddr* you can cast your struct sockaddr_in* to that type with ease and safety.
    if (bind(httpd, (struct sockaddr *)&name, sizeof(name)) <0 )
        error_die("bind")
    if (*port == 0) {
        socklen_t namelen = sizeof(name)
        // 获得名字然后写入 httpd
        if (getsockname(httpd, (struct sockaddr *)&name, &namelen) == -1 )
            error_die("getsockname");
        // 这两步的正反含义是啥?
        *port = ntohs(name.sin_port)
    }
    // 监听
    if (listen(httpd, 5) < 0 )
        error_die("listen")
    return(httpd)
}
/**
 * https://man7.org/linux/man-pages/man3/pthread_create.3.html
 * int pthread_create(pthread_t *restrict thread,
 *                        const pthread_attr_t *restrict attr,
 *                        void *(*start_routine)(void *),
 *                        void *restrict arg);
 *
 *
 */
int main(void)
{
    int server_sock = -1;
    u_short port = 4000;
    int client_sock = -1;
    struct sockaddr_in client_name;
    socklen_t client_name_len = sizeof(client_name)
    pthread_t newthread;

    server_sock = startup(&port);
    printf("httpd running on port %d\n", port);

    while(1)
    {
        // 监听
        client_sock = accept(server_sock,
            (struct sockaddr *)&client_name,
            &client_name_len);
        if (client_sock == -1)
            error_die("accept");
        // 创建一个线程处理请求
        if (pthread_create(&newthread, NULL,
            (void *)accept_request, (void *)(intptr_t)client_sock) != 0 )
            perror("pthread_create");
    }
    close(server_sock)
    return(0)
}

