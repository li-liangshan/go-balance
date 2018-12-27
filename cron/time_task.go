/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/27
 * @Last Modified by: liliangshan
 * *************************************************/

package cron

type timeTasks []*Task

func (tasks timeTasks) Len() int {
	return len(tasks)
}

func (tasks timeTasks) Swap(i, j int) {
	tasks[i], tasks[j] = tasks[j], tasks[i]
}

func (tasks timeTasks) Less(i, j int) bool {
	if tasks[i].Next.IsZero() {
		return false
	}
	if tasks[j].Next.IsZero() {
		return true
	}

	return tasks[i].Next.Before(tasks[j].Next)
}


