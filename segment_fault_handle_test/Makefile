#
# Makefile 
# 使用方法：
# 0.配置Makefile中的lib，inc等路径，以及编译选项
# 1.直接make(对应文件夹下的所有c文件都会编译)
# 2.部分编译：make src='test.c aad.c'
#	-- by jihan
#

DIR_INC = .
DIR_SRC = .
DIR_OBJ = .
DIR_BIN = .

DEFAULT_LIB = /usr/lib64/
#APR_INC = `pkg-config --cflags apr-1`
#G_INC = `pkg-config --cflags glib-2.0`
#REDIS_INC = /usr/local/include/hiredis/


ALL_LD = -L${DEFAULT_LIB} 
ALL_INC = 
ALL_LIB = 

CXX = g++
CC = gcc
CFLAGS = -g -rdynamic -Wall -Werror #-O2,不能和-rdynamic共存？
MPIF90 = mpiif90

SRC=$(src)
ifeq ("X${SRC}", "X")
SRC = $(wildcard ${DIR_SRC}/*.c)
OBJ = $(patsubst %.c,${DIR_OBJ}/%.o,$(notdir ${SRC}))
else
OBJ = $(patsubst %.c,${DIR_OBJ}/%.o,${SRC})
endif

TARGET = a.out

BIN_TARGET = ${DIR_BIN}/${TARGET}


.PHONY:all
all: ${BIN_TARGET}

${BIN_TARGET}: ${OBJ}
	$(CC) $(CFLAGS) $(OBJ)  -o $@ $(ALL_LD) $(ALL_LIB)

${DIR_OBJ}/%.o:${DIR_SRC}/%.c
	$(CC) $(CFLAGS)  -c $< -o $@ $(ALL_INC)

.PHONY:clean
clean:
	-rm -f *.out
	-rm -f *.o

.PHONY:maketest
maketest:
	@-echo "source:"${SRC}
	@-echo "object:"${OBJ} " target:"${BIN_TARGET}
	@-echo $(notdir ${SRC})

#（1）Makefile中的 符号 $@, $^, $< 等意思：
#　　$@  表示目标文件
#　　$^  表示所有的依赖文件
#　　$<  表示第一个依赖文件
#　　$?  表示比目标还要新的依赖文件列表
#	 - 	 表示出错继续执行
#	 @   表示命令不回显

#（2）wildcard、notdir、patsubst的意思：
#　　wildcard : 扩展通配符
#　　notdir ： 去除路径
#　　patsubst ：替换通配符

#（3）PHONY 表示伪目标，比如clean，如果文件中
#	  存在clean文件，那么添加了此关键字依旧能
#	  执行rm操作。


