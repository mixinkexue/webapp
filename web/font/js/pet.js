

function adoptAPet(id) {
    const formData = new FormData();
    formData.set('pet_id', id);
    fetch( '/webapp/pet/adopt', {
        method: "POST",
        body: formData,
        credentials: 'include',

    })
        .then(response => {
            if (response.ok) {
                return response.json()
            } else {
                throw new Error('后台服务故障');
            }
        })
        .then(data => {
                if (data.code === 200) {
                    // 修改成功，关闭弹出框并刷新页面
                    alert("领养成功，快到我的宠物看看吧");
                    location.reload();
                } else {
                    alert(data.data)
                }
            }
        )
        .catch(error => {
            alert(error);
        });

}