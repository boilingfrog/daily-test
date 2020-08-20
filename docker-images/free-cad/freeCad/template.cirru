
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
        div.banner-title
          = __INSERT_TITLE__
        div
      div.pad-area
        a.download-link
          :href __INSERT_DOWNLOAD__
          :target _blank
          button.download-button
            div.download-icon
            span
              = "下载应用"
    div.page-body
      div.v-space
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
      div.v-space
      @with (@ aggregation)
        div.content-section
          div.text-content
            div.title
              @ title
            div.description
              @ description
          div.content-space
          div.demo-image
      div.v-space
