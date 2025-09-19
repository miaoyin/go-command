import os
import datetime
import subprocess
import argparse

#参数
parser = argparse.ArgumentParser(description='bo build')
sub_argparse = parser.add_subparsers(title="子命令", dest="command", help="build, clean")
#编译
build_parser = sub_argparse.add_parser("build", help="编译go")
build_parser.add_argument('--version', dest='version', type=str, help='版本号 0.0.1')
build_parser.add_argument('--prefix', dest='prefix', type=str, help='前缀 demo')
build_parser.add_argument('--goos', dest='goos', type=str, help='操作系统linux, windows')
build_parser.add_argument('--goarch', dest='goarch', type=str, help='硬件架构 amd64, arm, arm64')
build_parser.add_argument('--gotags', dest='gotags', type=str, help='标签 plugin,imagelib')

#清理
clean_parser = sub_argparse.add_parser("clean", help="清理")
clean_parser.add_argument('--version', dest='version', type=str, help='版本号 0.0.1')
clean_parser.add_argument('--prefix', dest='prefix', type=str, help='前缀 demo')


def format_out_name(exe_prefix, version, commit_id, goos, goarch):
    """
    格式化输出名
    :param exe_prefix: 前缀 demo
    :param version:  版本号 0.0.1
    :param commit_id: git提交码
    :param goos:  操作系统linux, windows
    :param goarch: 硬件架构 amd64, arm, arm64
    :return:
    """
    ext = ".exe"
    if goos != "windows":
        ext = "."+goarch
    return f"{exe_prefix}-{version}-{commit_id}{ext}"

def format_ld_flags(version, commid_id, build_time):
    """
    # ldflags选项
    :param version:
    :param commid_id:
    :param build_time:
    :return:
    """
    return f"-ldflags \"-X 'main.Version={version}' -X 'main.CommitId={commid_id}' -X 'main.BuildTime={build_time}' -w -s -linkmode internal\" "

# 直接编译
def build(goos, goarch, out_name, go_ld_flags, go_tags):
    """
    :param goos: 操作系统
    :param goarch:  架构
    :param out_name: 输出名
    :param go_ld_flags: 嵌入flags
    :param go_tags: 标签, 条件编译
    :return:
    """
    print(f"==============build {goos} {goarch}===============")
    build_env = dict(CGO_ENABLED="0", GOOS=goos, GOARCH=goarch)
    command = f"go build -o {out_name} {go_ld_flags} {go_tags}"
    print(f"command: {command}")
    print(f"env: {build_env}")
    subprocess.run(command, env=dict(os.environ, **build_env))

def clean(exe_prefix, version):
    """
    清理
    :param exe_prefix:
    :param version:
    :return:
    """
    subprocess.run(f"rm -rf {exe_prefix}-{version}*")


def build_command(args):
    build_time = datetime.datetime.now().strftime("%Y-%m-%d %H:%M")
    commit_id = subprocess.getoutput("git rev-parse --short HEAD")
    out_name = format_out_name(args.prefix, args.version, commit_id, args.goos, args.goarch)
    ld_flags = format_ld_flags(args.version, commit_id, build_time)
    go_tags = f"--tags {args.gotags}" if args.gotags is not None else ""
    build(args.goos, args.goarch, out_name, ld_flags, go_tags)

def clean_command(args):
    os.system(f"rm -rf {args.prefix}-{args.version}*")


def main():
    # 参数
    args = parser.parse_args()
    print(f"args: {args}")
    if args.command=="build":
        # 编译
        build_command(args)
    if args.command=="clean":
        #clean
        clean_command(args)


if __name__ == "__main__":
    main()

