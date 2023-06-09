fn main() {
    let _ = sentry::init(("https://e0a450c5ab224248b6f4a471f061cc9c@o4505029175345152.ingest.sentry.io/4505121694416896", sentry::ClientOptions {
        release: sentry::release_name!(),
        ..Default::default()
    }));
    println!("Hello, world!");
}
