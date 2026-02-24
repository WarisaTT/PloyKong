import type { DirectiveBinding } from 'vue'

export const intersectDirective = {
  mounted(el: HTMLElement, binding: DirectiveBinding) {
    // If IntersectionObserver is not supported, just make it visible immediately
    if (!window.IntersectionObserver) {
      el.classList.add('is-visible')
      return
    }

    // Default options
    const options = {
      root: null,
      rootMargin: '0px 0px -50px 0px', // Trigger slightly before it comes into view
      threshold: 0.1, // Trigger when 10% of the element is visible
      ...((binding.value as object) || {})
    }

    const observer = new IntersectionObserver((entries, observerInstance) => {
      entries.forEach(entry => {
        if (entry.isIntersecting) {
          entry.target.classList.add('is-visible')
          // By default, animate only once
          observerInstance.unobserve(entry.target)
        }
      })
    }, options)

    observer.observe(el)
  }
}
