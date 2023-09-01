<template>
  <nav class="navbar navbar-expand-lg navbar-dark">
    <div class="container-fluid">
      <a class="navbar-brand" href="/">
        <div
          style="
            border: 0px solid white;
            border-radius: 4px;
            padding: 5px;
            background-color: white;
            background-image: url('/logo.png');
            background-size: cover;
            width: 30px;
            height: 30px;
          "
          class="d-inline-block align-top"
        ></div>
        <span style="font-size: 1.1rem" class="m-2">环境监测系统 API</span>
      </a>

      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarNav"
        aria-controls="navbarNav"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarNav">
        <div class="me-auto"></div>
        <ul class="navbar-nav">
          <li class="nav-item">
            <a class="nav-link" href="/">主页</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="/cross-station">数据对比</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="/contact-us">联系我们</a>
          </li>
        </ul>
        <form class="form-inline" style="margin-left: 10px">
          <a class="btn btn-outline-warning my-sm-0" type="button" href="/login"> 登入管理后台 </a>
        </form>
      </div>
    </div>
  </nav>
</template>

<style scoped lang="scss">
nav {
  background-color: $nav-background;
}

.admin {
  color: white;
  font-size: 1rem;
  font-weight: 400;
  margin-left: auto;
  margin-right: 1rem;
  padding: 0.5rem;
}
</style>
