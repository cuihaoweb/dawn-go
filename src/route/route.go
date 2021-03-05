package route

// Route 路由模块
type Route struct {
	GetContainer    *DataStructure
	PostContainer   *DataStructure
	DeleteContainer *DataStructure
	PutContainer    *DataStructure
}

// Get get请求
func (r *Route) Get(url string, handler handler) *Route {
	// add url and handler to GetContainer
	r.GetContainer.Add(url, handler)
	return r
}

// Post post请求
func (r *Route) Post(url string, handler handler) *Route {
	// add url and handler to PostContainer
	r.PostContainer.Add(url, handler)
	return r
}

// Delete delete请求
func (r *Route) Delete(url string, handler handler) *Route {
	// add url and handler to DeleteContainer
	r.DeleteContainer.Add(url, handler)
	return r
}

// Put put请求
func (r *Route) Put(url string, handler handler) *Route {
	// add url and handler to PutContainer
	r.PutContainer.Add(url, handler)
	return r
}
