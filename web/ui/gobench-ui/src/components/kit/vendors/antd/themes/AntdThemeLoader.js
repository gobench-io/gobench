/* eslint-disable */
const GetPostProcessor = function () {
  function Processor(themeName) {
    this.options = { wrappers: [`[data-kit-theme="${themeName}"]`] } || {}
  }

  Processor.prototype = {
    process: function (css) {
      // Check if css is wrapped with any one of the marks
      const antDesignTemplateMark = this.options.wrappers.reduce((mark, wrapper) => (css.indexOf(wrapper) >= 0 ? wrapper : mark), '')

      if (!antDesignTemplateMark) {
        return css
      }

      // Escape the mark to be RegExp friendly
      const escapedMark = antDesignTemplateMark.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')
      // Split, map, remove all, add one, clean double spaces
      const lines = css.split('\n').map((line) => {
        if (line.indexOf(antDesignTemplateMark) === -1) {
          return line
        }

        const finder = new RegExp(escapedMark, 'g')
        const finderNested = new RegExp(`(?<!^(?<!s))${escapedMark} `, 'g')
        const replaceNestedLine = line.replace(finderNested, '')
        const replaceLine = `${antDesignTemplateMark} ${replaceNestedLine.replace(finder, '')}`
        return replaceLine.replace(/( {2})/g, ' ')
      })

      return lines.join('\n')
    }
  }

  return Processor;
}

module.exports = {
  install: function (less, pluginManager, functions) {
    const PostProcessor = GetPostProcessor(less)
    functions.add('apply', function (theme) {
      pluginManager.addPostProcessor(new PostProcessor(theme.value))
      return false;
    });
  },
}
