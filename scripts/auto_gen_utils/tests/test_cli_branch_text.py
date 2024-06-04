import config


def test_cli_branch_text():
    generated_branch = "generated-auto-v2-public-JavaSDK-DEXREQ-991-2020-01-21-20-16-02"
    expected_branch = "generated-auto-v2-public-PythonCLI-DEXREQ-991-2020-01-21-20-16-02"
    cli_branch = get_cli_branch_text(generated_branch)
    assert expected_branch == cli_branch

    generated_branch = "generated-auto-v2-public-RubySDK-DEXREQ-991-2020-01-21-20-16-02"
    expected_branch = "generated-auto-v2-public-PythonCLI-DEXREQ-991-2020-01-21-20-16-02"
    cli_branch = get_cli_branch_text(generated_branch)
    assert expected_branch == cli_branch


def test_cli_branch_text_with_debug_prefix():
    generated_branch = "debug-generated-auto-v2-public-JavaSDK-DEXREQ-991-2020-01-21-20-16-02"
    expected_branch = "debug-generated-auto-v2-public-PythonCLI-DEXREQ-991-2020-01-21-20-16-02"
    cli_branch = get_cli_branch_text(generated_branch)
    assert expected_branch == cli_branch

    generated_branch = "debug-generated-auto-v2-public-RubySDK-DEXREQ-991-2020-01-21-20-16-02"
    expected_branch = "debug-generated-auto-v2-public-PythonCLI-DEXREQ-991-2020-01-21-20-16-02"
    cli_branch = get_cli_branch_text(generated_branch)
    assert expected_branch == cli_branch


def get_cli_branch_text(gen_branch):
    for tool in config.TOOL_NAMES:
        if tool in gen_branch:
            return gen_branch.replace(tool, config.CLI_NAME)
    return gen_branch
