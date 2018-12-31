/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/31
 * @Last Modified by: liliangshan
 * *************************************************/

/**
   • 类型 T ⽅方法集包含全部 receiver T ⽅方法。
   • 类型 *T ⽅方法集包含全部 receiver T + *T ⽅方法。
   • 如类型 S 包含匿名字段 T，则 S ⽅方法集包含 T ⽅方法。
	• 如类型 S 包含匿名字段 *T，则 S ⽅方法集包含 T + *T ⽅方法。
	• 不管嵌⼊入 T 或 *T，*S ⽅方法集总是包含 T + *T ⽅方法。
   根据调⽤用者不同，⽅方法分为两种表现形式:
		instance.method(args...) ---> <type>.func(instance, args...)
	   前者称为 method value，后者 method expression。
	   两者都可像普通函数那样赋值和传参，区别在于 method value 绑定实例，⽽而 method expression 则须显式传参。
 */

package items

import "fmt"

type Queue struct {
	elements []interface{}
}

func NewQueue() *Queue {
	return &Queue{make([] interface{}, 10)}
}

func (*Queue) Push(e interface{}) error {
	panic("not implements")
}

func (queue *Queue) Len() int {
	return len(queue.elements)
}

type Data struct {
	x int
}

func (data Data) ValueTest() {
	fmt.Printf("Value: %p\n", &data)
}

func (data *Data) PointerTest() {
	fmt.Printf("Pointer: %p\n", data)
}

type User struct {
	id   int
	name string
}

func (user *User) Test() {
	fmt.Printf("%p, %v\n", user, user)
}

func UserTest() {
	u := User{1, "Tom"}
	u.Test()

	mValue := u.Test
	mValue()

	mExpression := (*User).Test
	mExpression(&u)
}
