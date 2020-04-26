package models

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

type SshHost struct {
	Id int
	Host      string `orm:"size(32)"`
	Port      int
	Username  string `orm:"size(32)"`
	Password  string `orm:"size(64)"`
}

type SshResult struct {
	Id int
	Host    string
	Success string
	Type string
	Content string
	Result  string
	Time time.Time
}

func NewSshHost(h,u,p string,port int) *SshHost{
	return &SshHost{
	   Host:h,
	   Username:u,
	   Password:p,
	   Port:port,
	}
}

func AddSSHHost(s *SshHost) error{
	ormer := orm.NewOrm()
	_, e := ormer.Insert(s)
	if e!=nil{
		return e
	}
	return nil
}

func AddSshResult(s *SshResult) error{
	ormer := orm.NewOrm()
	_, e := ormer.Insert(s)
	if e!=nil{
		return e
	}
	return nil
}

func DeleteSshHost(s *SshHost){
	ormer := orm.NewOrm()
	ormer.Delete(s)
}

func GetSshHost(id int) *SshHost{
	sshhost := &SshHost{
		Id:id,
	}
    o :=orm.NewOrm()
	o.Read(sshhost,"Id")
	return sshhost
}

func GetAllSshHost() ([]*SshHost){
	var hosts []*SshHost
	o :=orm.NewOrm()
	o.QueryTable("SshHost").All(&hosts)
	return hosts
}

func GetAllSshResult() ([]*SshResult){
	var results []*SshResult
	o :=orm.NewOrm()
	o.QueryTable("SshResult").All(&results)
	return results
	
}

func Connect(h *SshHost) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		config       ssh.Config
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(h.Password))
	var signer ssh.Signer
	auth = append(auth, ssh.PublicKeys(signer))
	config = ssh.Config{
		Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
		}

	
	clientConfig = &ssh.ClientConfig{
		User:    h.Username,
		Auth:    auth,
		Timeout: 30 * time.Second,
		Config:  config,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	
	// connet to ssh
	addr = fmt.Sprintf("%s:%d", h.Host, h.Port)
	
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}
	
	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}
	
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		return nil, err
	}
	
	return session, nil
}

func ExecCommands(host *SshHost,session *ssh.Session,cmdlist []string) (ret *SshResult){
	ret = &SshResult{}
	ret.Host = host.Host
	
	defer session.Close()
	//获取ssh客户端的输入句柄
	stdinBuf, err := session.StdinPipe()
	if err != nil {
		ret.Success = "失败"
		return
	}
	
	var outbt, errbt bytes.Buffer
	//获取ssh客户端的标准输出
	session.Stdout = &outbt
	//获取ssh客户端的标准错误
	session.Stderr = &errbt
	err = session.Shell()
	if err != nil {
		ret.Success = "失败"
		return
	}
	//在命令后面追加exit
	cmdlist = append(cmdlist,"exit")
	for _, c := range cmdlist {
		c = c + "\n"
		stdinBuf.Write([]byte(c))
		
	}
	session.Wait()
	ret.Success = "成功"
	ret.Result = outbt.String() + errbt.String()
	ret.Time = time.Now()
	ret.Type = "命令执行"
	ret.Content = "ls"
	return
}





