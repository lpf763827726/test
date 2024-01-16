use tracing::error;

use crate::app_result::AppResult;

pub async fn create_or_update_project_languages() -> AppResult<()> {
    let project_ids = crate::remote_service::get_all_project_not_deleted().await?;
    let len = project_ids.len();
    // 限制每次的执行数
    let concurrency_limit = 100;
    let mut tasks = vec![];
    for i in (0..len).step_by(concurrency_limit) {
        let end = if i + concurrency_limit < len {
            i + concurrency_limit
        } else {
            len
        };
        let mut fs = vec![];
        for j in i..end {
            let input_clone = project_ids[j];
            let task =
                tokio::task::spawn(
                    async move { crate::git::file::get_languages_by_project_by_tokei(&input_clone).await.unwrap() },
                );
            fs.push(task);
        }
        let result = futures::future::try_join_all(fs).await.unwrap();
        let mut send_vec = vec![];
        for r in result {
            send_vec.push(r);
            if send_vec.len() == 20 {
                let task = tokio::task::spawn(async move {
                    let send_vec_clone = send_vec.clone();
                    match crate::remote_service::post_project_languages(send_vec_clone).await {
                        Ok(_) => {}
                        Err(e) => {
                            error!("post project language error: {:?}", e)
                        }
                    }
                });
                tasks.push(task);
                send_vec = vec![];
            }
        }
    }
    let result = futures::future::try_join_all(tasks).await.unwrap();
    Ok(())
}
