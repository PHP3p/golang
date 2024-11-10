git config --local --list
git config --local --unset user.name
git remote set-url origin git@github.com:PHP3p/your-repo.git
来测试你的 SSH 连接-- 秘钥加入 known_hosts
ssh -T git@github.com

