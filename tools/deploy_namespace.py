import os
import utils
import argparse
import yaml

parser = argparse.ArgumentParser()
parser.add_argument('--namespace', type=str, default='assisted-installer')
parser.add_argument("--deploy-namespace", type=lambda x: (str(x).lower() == 'true'), default=True)
args = parser.parse_args()


def main():
    if args.deploy_namespace is False:
        print("Not deploying namespace")
        return
    src_file = os.path.join(os.getcwd(), "deploy/namespace/namespace.yaml")
    dst_file = os.path.join(os.getcwd(), "build/namespace.yaml")
    with open(src_file, "r") as src:
        with open(dst_file, "w+") as dst:
            data = yaml.safe_load(src)
            data['metadata']['name'] = args.namespace
            data['metadata']['labels']['name'] = args.namespace
            print("Deploying {}".format(dst_file))
            yaml.dump(data, dst, default_flow_style=False)

    utils.apply(dst_file)


if __name__ == "__main__":
    main()
