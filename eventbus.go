// EventBus - Event Bus for SiYuan.
// Copyright (c) 2022-present, b3log.org
//
// EventBus is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
//
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//
// See the Mulan PSL v2 for more details.

package eventbus

import "github.com/asaskevich/EventBus"

var bus = EventBus.New()

func Publish(topic string, arg ...interface{}) {
	bus.Publish(topic, arg...)
}

func Subscribe(topic string, handler func()) error {
	return bus.Subscribe(topic, handler)
}
