
doctype

html
  head
    title
      = __INSERT_TITLE__
    link
      :rel stylesheet
      :href main.css?v=1
    meta (:name viewport)
      :content "width=device-width, initial-scale=1, maximum-scale=1, minimum-scale=1, user-scalable=no, viewport-fit=cover"
    meta
      :http-equiv Content-type
      :content "text/html; charset=utf-8"
  body
    div.page-head
      div.pad-area
      div
        div.title-container
          div.shadow-title $ @ varTitle
          div.banner-title  $ @ varTitle
          div.placeholder-title $ @ varTitle
        div.fake-slim-line
        div.fake-underline
      div.pad-area
        a.download-link
          :href $ @ varLink
          :target _blank
          button.download-button
            div.download-icon
            span
              @ downloadApp
    div.page-body
      @with (@ platform)
        div.content-section
          div.demo-image
          div.content-space
          div.text-content
            div.title
              @ title
            div.sub-title
              @ subTitle
            div.description
              @ description

      @with (@ aggregation)
        div.content-section
          div.text-content
            div.title
              @ title
            div.description
              @ description
          div.content-space
          div.demo-image
