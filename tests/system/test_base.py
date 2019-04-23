from bb8beat import BaseTest

import os


class Test(BaseTest):

    def test_base(self):
        """
        Basic test with exiting Bb8beat normally
        """
        self.render_config_template(
            path=os.path.abspath(self.working_dir) + "/log/*"
        )

        bb8beat_proc = self.start_beat()
        self.wait_until(lambda: self.log_contains("bb8beat is running"))
        exit_code = bb8beat_proc.kill_and_wait()
        assert exit_code == 0
