let dom;
if (!dom) {
    dom = {
        element(selector) {
            const s = dom.selectable(this)
            // return s.querySelector(selector)
            dom.html = s.querySelector(selector).innerHTML
            ipc.emit('implantName', [s.querySelector(selector).innerHTML]);
        },

        elements(selector) {
            // return dom.selectable(this).querySelectorAll(selector)
            let data = []
            let doms = dom.selectable(this).querySelectorAll(selector)
            doms.forEach(function (dom) {
                data.push(dom.innerHTML)
            })
            dom.html = data
        },

        elementX(xPath) {
            const s = dom.selectable(this)
            // return document.evaluate(xPath, s, null, XPathResult.FIRST_ORDERED_NODE_TYPE).singleNodeValue
            let el = document.evaluate(xPath, s, null, XPathResult.FIRST_ORDERED_NODE_TYPE).singleNodeValue
            if (el) {
                dom.html = el.innerHTML
            } else {
                dom.html = ""
            }
        },

        elementsX(xpath) {
            const s = dom.selectable(this)
            const iter = document.evaluate(xpath, s, null, XPathResult.ORDERED_NODE_ITERATOR_TYPE)
            const list = []
            let el
            while ((el = iter.iterateNext())) list.push(el.innerHTML)
            // return list
            dom.html = list
        },
        
        selectable(s) {
            return s.querySelector ? s : document
        },

        elementR(selector, regex) {
            var reg
            var m = regex.match(/(\/?)(.+)\1([a-z]*)/i)
            // cSpell:ignore gmix
            if (m[3] && !/^(?!.*?(.).*?\1)[gmixXsuUAJ]+$/.test(m[3]))
                reg = new RegExp(regex)
            else reg = new RegExp(m[2], m[3])

            const s = dom.selectable(this)
            const el = Array.from(s.querySelectorAll(selector)).find((e) =>
                reg.test(dom.text.call(e))
            )
            // return el ? el : null
            dom.html = el ? el.innerHTML : ""
        }
    }
}